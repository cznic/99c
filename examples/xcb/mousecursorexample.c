// 99c mousecursorexample.c -lxcb && ./a.out

// +build ignore

// src: https://www.x.org/releases/current/doc/libxcb/tutorial/index.html#mousecursorexample

#include <stdlib.h>
#include <stdio.h>
#include <string.h>

#include <xcb/xcb.h>

#define WIDTH 300
#define HEIGHT 150

static xcb_gc_t gc_font_get(xcb_connection_t * c, xcb_screen_t * screen, xcb_window_t window, const char *font_name);

static void button_draw(xcb_connection_t * c, xcb_screen_t * screen, xcb_window_t window, int16_t x1, int16_t y1, const char *label);

static void text_draw(xcb_connection_t * c, xcb_screen_t * screen, xcb_window_t window, int16_t x1, int16_t y1, const char *label);

static void cursor_set(xcb_connection_t * c, xcb_screen_t * screen, xcb_window_t window, int cursor_id);

static void button_draw(xcb_connection_t * c, xcb_screen_t * screen, xcb_window_t window, int16_t x1, int16_t y1, const char *label)
{
	xcb_point_t points[5];
	xcb_void_cookie_t cookie_gc;
	xcb_void_cookie_t cookie_line;
	xcb_void_cookie_t cookie_text;
	xcb_generic_error_t *error;
	xcb_gcontext_t gc;
	int16_t width;
	int16_t height;
	uint8_t length;
	int16_t inset;

	length = strlen(label);
	inset = 2;

	gc = gc_font_get(c, screen, window, "7x13");

	width = 7 * length + 2 * (inset + 1);
	height = 13 + 2 * (inset + 1);
	points[0].x = x1;
	points[0].y = y1;
	points[1].x = x1 + width;
	points[1].y = y1;
	points[2].x = x1 + width;
	points[2].y = y1 - height;
	points[3].x = x1;
	points[3].y = y1 - height;
	points[4].x = x1;
	points[4].y = y1;
	cookie_line = xcb_poly_line_checked(c, XCB_COORD_MODE_ORIGIN, window, gc, 5, points);

	error = xcb_request_check(c, cookie_line);
	if (error) {
		fprintf(stderr, "ERROR: can't draw lines : %d\n", error->error_code);
		xcb_disconnect(c);
		exit(-1);
	}

	cookie_text = xcb_image_text_8_checked(c, length, window, gc, x1 + inset + 1, y1 - inset - 1, label);
	error = xcb_request_check(c, cookie_text);
	if (error) {
		fprintf(stderr, "ERROR: can't paste text : %d\n", error->error_code);
		xcb_disconnect(c);
		exit(-1);
	}

	cookie_gc = xcb_free_gc(c, gc);
	error = xcb_request_check(c, cookie_gc);
	if (error) {
		fprintf(stderr, "ERROR: can't free gc : %d\n", error->error_code);
		xcb_disconnect(c);
		exit(-1);
	}
}

static void text_draw(xcb_connection_t * c, xcb_screen_t * screen, xcb_window_t window, int16_t x1, int16_t y1, const char *label)
{
	xcb_void_cookie_t cookie_gc;
	xcb_void_cookie_t cookie_text;
	xcb_generic_error_t *error;
	xcb_gcontext_t gc;
	uint8_t length;

	length = strlen(label);

	gc = gc_font_get(c, screen, window, "7x13");

	cookie_text = xcb_image_text_8_checked(c, length, window, gc, x1, y1, label);
	error = xcb_request_check(c, cookie_text);
	if (error) {
		fprintf(stderr, "ERROR: can't paste text : %d\n", error->error_code);
		xcb_disconnect(c);
		exit(-1);
	}

	cookie_gc = xcb_free_gc(c, gc);
	error = xcb_request_check(c, cookie_gc);
	if (error) {
		fprintf(stderr, "ERROR: can't free gc : %d\n", error->error_code);
		xcb_disconnect(c);
		exit(-1);
	}
}

static xcb_gc_t gc_font_get(xcb_connection_t * c, xcb_screen_t * screen, xcb_window_t window, const char *font_name)
{
	uint32_t value_list[3];
	xcb_void_cookie_t cookie_font;
	xcb_void_cookie_t cookie_gc;
	xcb_generic_error_t *error;
	xcb_font_t font;
	xcb_gcontext_t gc;
	uint32_t mask;

	font = xcb_generate_id(c);
	cookie_font = xcb_open_font_checked(c, font, strlen(font_name), font_name);

	error = xcb_request_check(c, cookie_font);
	if (error) {
		fprintf(stderr, "ERROR: can't open font : %d\n", error->error_code);
		xcb_disconnect(c);
		return -1;
	}

	gc = xcb_generate_id(c);
	mask = XCB_GC_FOREGROUND | XCB_GC_BACKGROUND | XCB_GC_FONT;
	value_list[0] = screen->black_pixel;
	value_list[1] = screen->white_pixel;
	value_list[2] = font;
	cookie_gc = xcb_create_gc_checked(c, gc, window, mask, value_list);
	error = xcb_request_check(c, cookie_gc);
	if (error) {
		fprintf(stderr, "ERROR: can't create gc : %d\n", error->error_code);
		xcb_disconnect(c);
		exit(-1);
	}

	cookie_font = xcb_close_font_checked(c, font);
	error = xcb_request_check(c, cookie_font);
	if (error) {
		fprintf(stderr, "ERROR: can't close font : %d\n", error->error_code);
		xcb_disconnect(c);
		exit(-1);
	}

	return gc;
}

static void cursor_set(xcb_connection_t * c, xcb_screen_t * screen, xcb_window_t window, int cursor_id)
{
	uint32_t values_list[3];
	xcb_void_cookie_t cookie_font;
	xcb_void_cookie_t cookie_gc;
	xcb_generic_error_t *error;
	xcb_font_t font;
	xcb_cursor_t cursor;
	xcb_gcontext_t gc;
	uint32_t mask;
	uint32_t value_list;

	font = xcb_generate_id(c);
	cookie_font = xcb_open_font_checked(c, font, strlen("cursor"), "cursor");
	error = xcb_request_check(c, cookie_font);
	if (error) {
		fprintf(stderr, "ERROR: can't open font : %d\n", error->error_code);
		xcb_disconnect(c);
		exit(-1);
	}

	cursor = xcb_generate_id(c);
	xcb_create_glyph_cursor(c, cursor, font, font, cursor_id, cursor_id + 1, 0, 0, 0, 0, 0, 0);

	gc = xcb_generate_id(c);
	mask = XCB_GC_FOREGROUND | XCB_GC_BACKGROUND | XCB_GC_FONT;
	values_list[0] = screen->black_pixel;
	values_list[1] = screen->white_pixel;
	values_list[2] = font;
	cookie_gc = xcb_create_gc_checked(c, gc, window, mask, values_list);
	error = xcb_request_check(c, cookie_gc);
	if (error) {
		fprintf(stderr, "ERROR: can't create gc : %d\n", error->error_code);
		xcb_disconnect(c);
		exit(-1);
	}

	mask = XCB_CW_CURSOR;
	value_list = cursor;
	xcb_change_window_attributes(c, window, mask, &value_list);

	xcb_free_cursor(c, cursor);

	cookie_font = xcb_close_font_checked(c, font);
	error = xcb_request_check(c, cookie_font);
	if (error) {
		fprintf(stderr, "ERROR: can't close font : %d\n", error->error_code);
		xcb_disconnect(c);
		exit(-1);
	}
}

int main()
{
	xcb_screen_iterator_t screen_iter;
	xcb_connection_t *c;
	const xcb_setup_t *setup;
	xcb_screen_t *screen;
	xcb_generic_event_t *e;
	xcb_generic_error_t *error;
	xcb_void_cookie_t cookie_window;
	xcb_void_cookie_t cookie_map;
	xcb_window_t window;
	uint32_t mask;
	uint32_t values[2];
	int screen_number;
	uint8_t is_hand = 0;

	/* getting the connection */
	c = xcb_connect(NULL, &screen_number);
	if (!c) {
		fprintf(stderr, "ERROR: can't connect to an X server\n");
		return -1;
	}

	/* getting the current screen */
	setup = xcb_get_setup(c);

	screen = NULL;
	screen_iter = xcb_setup_roots_iterator(setup);
	for (; screen_iter.rem != 0; --screen_number, xcb_screen_next(&screen_iter))
		if (screen_number == 0) {
			screen = screen_iter.data;
			break;
		}
	if (!screen) {
		fprintf(stderr, "ERROR: can't get the current screen\n");
		xcb_disconnect(c);
		return -1;
	}

	/* creating the window */
	window = xcb_generate_id(c);
	mask = XCB_CW_BACK_PIXEL | XCB_CW_EVENT_MASK;
	values[0] = screen->white_pixel;
	values[1] = XCB_EVENT_MASK_KEY_RELEASE | XCB_EVENT_MASK_BUTTON_PRESS | XCB_EVENT_MASK_EXPOSURE | XCB_EVENT_MASK_POINTER_MOTION;
	cookie_window = xcb_create_window_checked(c, screen->root_depth, window, screen->root, 20, 200, WIDTH, HEIGHT, 0, XCB_WINDOW_CLASS_INPUT_OUTPUT, screen->root_visual, mask, values);
	cookie_map = xcb_map_window_checked(c, window);

	/* error managing */
	error = xcb_request_check(c, cookie_window);
	if (error) {
		fprintf(stderr, "ERROR: can't create window : %d\n", error->error_code);
		xcb_disconnect(c);
		return -1;
	}
	error = xcb_request_check(c, cookie_map);
	if (error) {
		fprintf(stderr, "ERROR: can't map window : %d\n", error->error_code);
		xcb_disconnect(c);
		return -1;
	}

	cursor_set(c, screen, window, 68);

	xcb_flush(c);

	while (1) {
		e = xcb_poll_for_event(c);
		if (e) {
			switch (e->response_type & ~0x80) {
			case XCB_EXPOSE:{
					char *text;

					text = "click here to change cursor";
					button_draw(c, screen, window, (WIDTH - 7 * strlen(text)) / 2, (HEIGHT - 16) / 2, text);

					text = "Press ESC key to exit...";
					text_draw(c, screen, window, 10, HEIGHT - 10, text);
					break;
				}
			case XCB_BUTTON_PRESS:{
					xcb_button_press_event_t *ev;
					int length;

					ev = (xcb_button_press_event_t *) e;
					length = strlen("click here to change cursor");

					if ((ev->event_x >= (WIDTH - 7 * length) / 2) && (ev->event_x <= ((WIDTH - 7 * length) / 2 + 7 * length + 6)) && (ev->event_y >= (HEIGHT - 16) / 2 - 19) && (ev->event_y <= ((HEIGHT - 16) / 2)))
						is_hand = 1 - is_hand;

					is_hand ? cursor_set(c, screen, window, 58) : cursor_set(c, screen, window, 68);
				}
			case XCB_KEY_RELEASE:{
					xcb_key_release_event_t *ev;

					ev = (xcb_key_release_event_t *) e;

					switch (ev->detail) {
						/* ESC */
					case 9:
						free(e);
						xcb_disconnect(c);
						return 0;
					}
				}
			}
			free(e);
		}
	}

	return 0;
}
