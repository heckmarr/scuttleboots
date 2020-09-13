#!/bin/bash python3

import curses
stdscr = curses.initscr()
#Start curses

class Painter(object):
    def DrawScreen(self):
        #draw the screen here

    def Announce(self):
        #Announce to the back end that
        #The image is ready
    def ImgToAscii(self):
        #this function used to be called asciified

curses.endwin()
