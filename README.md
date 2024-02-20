# Slot-Machine

Go technical assessment for Massive Studios (Robert Tuohy)

## Command Arguments

Apply the following arguments when running:

| Flag | Argument | Description 
| ----------- | ----------- | -----------
| reelCount | String | Number of reels the application will create/run
| reelSymbolLength | String | Number of symbols within the reels themselves (eg. ["A", "K", "Q", "J"] would equal 4.)

Then add the Symbols themselves (space separated) followed by their respective Multiplier values

So to run the application as intended for this task.

`go run main.go -reelCount=4 -reelSymbolLength=6 A K Q J 10 X 20 15 10 5 2 0`

Flag values are converted to integers once within the app itself.

## Structure

I created 4 reels of equal length and equal in symbols for this particular game:

["A", "K", "Q", "J", "10", "X"]

["A", "K", "Q", "J", "10", "X"]

["A", "K", "Q", "J", "10", "X"]

["A", "K", "Q", "J", "10", "X"]

## Answering questions posed within the task PDF

**Can the application inject a different implementation of a Game Engine and still works?**

Partially, you could add in different symbols and amount of symbols to get a similar game put with slightly different parameters.

However if I had more time I would try to generalise this even further. Maybe the engine has a number of different running modes and the start script/flags itself upon startup would describe which 'mode' the app would run in and act accordingly.


**Can the Game Engine handle different winning conditions, mechanisms and pay tables?**

It can handle different pay tables by altering the start script. 

Again, given more time I would again try to understand the different winning conditions and mechanisms that exist in this industry, implement them all within the code. Then again through a start script type flag the program would know which winning conditions or mechanisms it should use. Allowing for a template to exist for many types of games within the same repo. 

Or at the very least, big chunks of code are tranferable to other projects, cutting down dev time and increasing chances of Massive Studios reaching their desired number of games out to market in 2024.

**Are you happy with the test coverage?**

No, I only added some unit tests for the helper functions. Would like to have built out more end to end tests. As stated in my initial interview i'm a little rusty so a good amount of time of me working on this task was spent trying to "grease the wheels" and get back into a go headspace and complete the primary task (the game itself). 

**Can you think of possible bug scenarios or improvements? If so feel free to include in the readme.md.**

1. Yeah I think some better validation on inputs/outputs at various points as there may be some bugs there. 

2. I think I could use some more go routines to utlise concurrency and optimise this app in certain areas. Maybe when building out the strings for console (would need to profile for validation of it being indeed quicker) for example. Didn't use channels within my implementation, because I knew the memory location (slice position) for each go func() I didn't have to worry about it. With greather exappnsion of feature logic would likely be some use cases for it. Maybe if there was multiple winning conditions they could be calculated concurrently rather than sequentially and winning values accumulated at the end.

3. I would have liked to do some profiling of what is quicker, strings.Join() or strings.Builder() for example. When creating the strings to output to the console (or eventual UI/Game if it was running as a live service). Went from memory that strings.Join() was quicker.

4. As stated before, better test coverage across the board.

5. Unsure whether my usage of one central struct and methods on that struct were the most optimal design choice (again, feel rusty on best practices etc. in Go).

6. I would want a better startup mechanism with flags etc, I dislike my current usage of...

    `go run main.go -reelCount=4 -reelSymbolLength=6 A K Q J 10 X 20 15 10 5 2 0`

    I think it's currently an ugly implementation (especially if reel symbols for each reel were going to vary, that list of spaced attributes at the end would be shockingly awful to look at, manage and update), but the general theme is what I'm going for. Which is an eventual "generic system" that can take arguments through something like a shell script (e.g. startup.sh). Then based off those flags/parameters passed in, the game can run/calculate/display accordingly.

    A better way is surely possible with some research, I didn't want to waste too much time on this initially as I thought the broader task and logic was more important to focus on in the few hours you state to work on it. 

7. I would like to add logic and tests for the parameters you talked about in the interview (RTP - return to player). Not that I have the greatest understanding of what a company would expect in those. Ido recall you may have stated 97%. Would implement test frameworks to spin N number of times to check everything is within expected bounds. Spins in values of million (1e6), billion(1e9) etc.

8. Add Logging generally, for example
- Errors/warnings/info.
- Usage stats/metrics/graphs (frequency, time between "spins") etc. 
- Amounts won/lost (from business persective).

9. Not exactly thrilled with my main fuction just looping, currently no nice way to break out. Now if user was to click the "X" (Exit) button and that sent something in as a sign to finish up like through an API for example. Then gather the necessary session stats, send those to a database and close gracefully. That would have been preferrable.

10. Get tests to run on startup and fail/alert if something went wrong

11. App currently running through `go run main.go` + flags etc.   This of course would eventually be built into a executable binary file I imagine, so need to build that and determine how to run that with flags parameters or bundle those values in as a build of game "X" and a different set of parameters could be game "Y"

**Bonus: Can different patterns be considered, other than the middle horizontal line?**

Yeah absolutely, can I just focused on the one stated in the task.

I did jump on your website and try a few of your games out. Some things would win in a diagonal for example.

A A A X

A A X A

A X A A

X A A A

For example the above pattern may also be a winner (X's in a diagonal pattern). I would define these patterns within the program, ascribe them set names, implement the checking logic and again add them in as flags upon startup to determine whether the reels should be checked in a different or multitude of different ways to ascertain whether the spin is a winning spin and calculate accordingly.



## Thanks

First off, thank you for the opportunity. Irrespective of how I do I've learnt a lot and gave me a very interesting task to dig into and get back doing some Golang. Very much enjoyed it. I was feeling very rusty at Go and its been great to get stuck into a set of interesting problems (lots of looking at the Go docs as I was frequently like "I know what i'm trying to do in my head, what's the idiomatic way to do this within the Go paradigm again).

I would very much appreciate some feedback on how i've done, the good and the bad. As I'm always looking to improve and learn better ways to plan, layout code and solve problems.

Look forward to chatting again

Rob
