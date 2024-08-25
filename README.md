# BeatSyncSolution

This is a personal project that I am making for me to properly download and manage all my music across multiple music streaming platforms (Spotify, SoundCloud, YouTube...)

I found that downloading all the songs one by one was very annoying, and I would generally have to write scripts to do it anyway. So I figured it would make more sense for me to make my own web application that will allow me to do this in whatever device I need it to. The main plan with this is to be able to access it, log into Spotify/SoundCloud or any of the other platforms, be able to see the track and playlists and say to directly download them. In the event that you are downloading a track from SoundCloud and theres a higher quality version found on Spotify, this will give you the option to get the Spotify version instead.

## Plan

* Application will be primarily written in Golang
* Application will be run using a Docker image
* Eventually BSS will be hosted in AWS
* Currently using existing projects (such as Zotify) to work on downloading the songs from each platform
