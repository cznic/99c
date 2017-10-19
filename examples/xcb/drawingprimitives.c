// 99c drawingprimitives.c -lxcb && ./a.out

// +build ignore

// src: https://www.x.org/releases/current/doc/libxcb/tutorial/index.html#drawingprim

#include <stdlib.h>
#include <stdio.h>

#include <xcb/xcb.h>

int main()
{
	xcb_connection_t *c;
	xcb_screen_t *screen;
	xcb_drawable_t win;
	xcb_gcontext_t foreground;
	xcb_generic_event_t *e;
	uint32_t mask = 0;
	uint32_t values[2];

	/* geometric objects */
	xcb_point_t points[] = {
		{10, 10},
		{10, 20},
		{20, 10},
		{20, 20}
	};

	xcb_point_t polyline[] = {
		{50, 10},
		{5, 20},	/* rest of points are relative */
		{25, -20},
		{10, 10}
	};

	xcb_segment_t segments[] = {
		{100, 10, 140, 30},
		{110, 25, 130, 60}
	};

	xcb_rectangle_t rectangles[] = {
		{10, 50, 40, 20},
		{80, 50, 10, 40}
	};

	xcb_arc_t arcs[] = {
		{10, 100, 60, 40, 0, 90 << 6},
		{90, 100, 55, 40, 0, 270 << 6}
	};

	/* Open the connection to the X server */
	c = xcb_connect(NULL, NULL);

	/* Get the first screen */
	screen = xcb_setup_roots_iterator(xcb_get_setup(c)).data;

	/* Create black (foreground) graphic context */
	win = screen->root;

	foreground = xcb_generate_id(c);
	mask = XCB_GC_FOREGROUND | XCB_GC_GRAPHICS_EXPOSURES;
	values[0] = screen->black_pixel;
	values[1] = 0;
	xcb_create_gc(c, foreground, win, mask, values);

	/* Ask for our window's Id */
	win = xcb_generate_id(c);

	/* Create the window */
	mask = XCB_CW_BACK_PIXEL | XCB_CW_EVENT_MASK;
	values[0] = screen->white_pixel;
	values[1] = XCB_EVENT_MASK_EXPOSURE;
	xcb_create_window(c,	/* Connection          */
			  XCB_COPY_FROM_PARENT,	/* depth               */
			  win,	/* window Id           */
			  screen->root,	/* parent window       */
			  0, 0,	/* x, y                */
			  150, 150,	/* width, height       */
			  10,	/* border_width        */
			  XCB_WINDOW_CLASS_INPUT_OUTPUT,	/* class               */
			  screen->root_visual,	/* visual              */
			  mask, values);	/* masks */

	/* Map the window on the screen */
	xcb_map_window(c, win);

	/* We flush the request */
	xcb_flush(c);

	while ((e = xcb_wait_for_event(c))) {
		switch (e->response_type & ~0x80) {
		case XCB_EXPOSE:{
				/* We draw the points */
				xcb_poly_point(c, XCB_COORD_MODE_ORIGIN, win, foreground, 4, points);

				/* We draw the polygonal line */
				xcb_poly_line(c, XCB_COORD_MODE_PREVIOUS, win, foreground, 4, polyline);

				/* We draw the segements */
				xcb_poly_segment(c, win, foreground, 2, segments);

				/* We draw the rectangles */
				xcb_poly_rectangle(c, win, foreground, 2, rectangles);

				/* We draw the arcs */
				xcb_poly_arc(c, win, foreground, 2, arcs);

				/* We flush the request */
				xcb_flush(c);

				break;
			}
		default:{
				/* Unknown event type, ignore it */
				break;
			}
		}
		/* Free the Generic Event */
		free(e);
	}

	return 0;
}
