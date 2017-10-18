// +build ignore

// src: https://www.x.org/releases/current/doc/libxcb/tutorial/index.html#helloworld

#include <stdio.h>
#include <unistd.h>		/* pause() */

#include <xcb/xcb.h>

int main()
{
	xcb_connection_t *c;
	xcb_screen_t *screen;
	xcb_window_t win;

	/* Open the connection to the X server */
	c = xcb_connect(NULL, NULL);

	/* Get the first screen */
	xcb_screen_iterator_t iter = xcb_setup_roots_iterator(xcb_get_setup(c));	//TODO bug
	screen = iter.data;

	/* Ask for our window's Id */
	win = xcb_generate_id(c);

	/* Create the window */
	xcb_create_window(c,	/* Connection          */
			  XCB_COPY_FROM_PARENT,	/* depth (same as root) */
			  win,	/* window Id           */
			  screen->root,	/* parent window       */
			  0, 0,	/* x, y                */
			  150, 150,	/* width, height       */
			  10,	/* border_width        */
			  XCB_WINDOW_CLASS_INPUT_OUTPUT,	/* class               */
			  screen->root_visual,	/* visual              */
			  0, NULL);	/* masks, not used yet */

	/* Map the window on the screen */
	xcb_map_window(c, win);

	/* Make sure commands are sent before we pause, so window is shown */
	xcb_flush(c);

	printf("Close the demo window and/or press ctrl-c while the terminal is focused to exit.\n");
	int i = pause();	/* hold client until Ctrl-C */
	printf("pause() returned %i\n", i);

	return 0;
}
