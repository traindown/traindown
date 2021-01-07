Traindown
=========

**A language to help athletes express their training. Inspired by [Markdown.](https://daringfireball.net/projects/markdown/)**

Traindown was created to help you easily capture all the things you do in a training session. It aims to provide you with a flexible, easy to remember language for expressing your training for record keeping purposes. As you will learn, writing out your training in **Traindown** is as simple as using a hashtag.

## The Basics

The minimum viable **Traindown** file could be expressed as
```
Squat: 500
```
This is saying "I did **one** rep of **500** in the **Squat**." 500 of what, you may ask? Well, that brings us to the next section!

## Metadata

**Traindown** provides tools for recording units (and any other data you can imagine) with **metadata**. Metadata are just **key value pairs** that begin with a **#**. You can have as much or as little metadata as you prefer, but the more the better in my opinion!

Let's add some metadata to the session from above to help us understand more about what happened that day.
```
# Unit: lbs
Squat: 500
```
Ah, that's better. So this session used pounds for the unit of weight. That squat was 500 _pounds_.

Metadata allows you to express facts about the context of your training. Some common keys are duration, day/week of training cycle, and if you have any injuries you are working through. Some folks capture things like their weight that day or even what they eat leading into it. Let's go ahead and some more data to our example session.
```
# Week: 2
# Day: 4
# Burritos: 3
# Unit: lbs

Squat: 500
```
All the metadata we just added to the top of the file is called **session metadata**. This information applies to the entire session. **Traindown** also supports the concept of **performance metadata**.

Each combination of a load for sets and reps is considered a **performance** of a given movement. So in our example, lets assume this trainee likes to capture their "reps in reserve" for their top sets.
```
# Week: 2
# Day: 4
# Burritos: 3
# Unit: lbs

Squat: 500
  # RIR: 2
```
This is saying "For the single squat at 500, I had 2 reps in reserve." Unlike the burrito count, the reps in reserve are applied to only the performance.

You may also track **movement metadata** that applies to all performances of the movement. To do so, just place your key value pair after the movement name and before any performances.

By adding more details in the metadata, you can make more sense of the _why_ in your training at both the session and perofmrance level. Metadata is also very searchable, so if you choose to use a **Traindown** enabled analysis tool, you can really start to make sense of the big picture.

But wait...one key pieces is missing. _When_ did this happen?

## Date and Time

The date and time when your training occurs is similar to metadata but is special enough to warrant it's own symbol. To specify the date and/or time that you trained, just use **@**, typically as the first thing in your **Traindown** file. To build on our example, it would look like
```
@ Dec 25 2019 8:03am

# Week: 2
# Day: 4
# Burritos: 3
# Unit: lbs

Squat: 500
  # RIR: 2
```
The format for the date and time is entirely up to you. Keep in mind if you are using a **Traindown** analytics tool, there may be a preferred format or formats, so be sure to read up on your particular tool.

## Sets and Reps

Now, our example is pretty unrealistic. Cold squatting 500 pounds, three burritos or none, is not common. Let's assume this is a more mortal trainee and fill in some of the details of the day.
```
@ Dec 25 2019 8:03am

# Week: 2
# Day: 4
# Burritos: 3
# Unit: lbs

Squat: 135 10r 2s 225 5r 315 5r 405 2r 455 500 3s
```
That's more realistic. So what was added here? It should read intuitively but to be sure, let's dive into the details.

Looking at
```
135 10r 2s
```
we can read this as "I did two **sets** of ten **reps** at 135 pounds on the Squat." As you can see, we denote sets with an **s** and reps with an **r**.

Further on down the day, we can see
```
225 5r
```
which can be read as "I did one set of five reps at 225 pounds on the Squat." If we omit either an r or and s, the default is 1. So circling back to the original 500, it _could_ be expressed as

**DON'T DO THIS**
```
500 1r 1s
```
**DON'T DO THIS**

_Note: Seriously, don't do this. Save yourself the 1r / 1s shenanigans._

but the whole point of **Traindown** is to make recording your training easy, not annoying.

The order of the reps and sets does not matter—we'll leave that debate to be settled on another date. However, you should _always_ lead a performance with the load.

There is one piece missing with our story so far. It's a dark thing. Not to be spoken of. That's right, I am talking about missed lifts.

If you are training and need to capture failed attempts, **Traindown** allows you to do that with the "failed rep" symbol **f**.

To keep our burrito powered example free of the scourage known as missed lifts, let's use a simple example to illustrate
```
Snatch:
  60 3s
  80 2s
  100
  110
  120 1f
  120
```
Here we can see that the lifter was successful up until the 120 at which point they happened to miss the first attempt there. They appear to have been able to correct and rally to successfully hoist that 120.

The **failed rep** allows you to capture the deficit in what you originally planned on hitting when used in conjunction with the rep symbol. This can be used to capture multiple failed reps, as well. For example
```
Glute bench raise:
  1000 10r
  1000 10r 6f
```
In the literal "butt blaster" example above, the trainee would be credited with 4 completed reps and 6 missed reps. The reason we use **10r** is to indicate the _reps attempted_. In the first example, the implied **1r** is considered and the trainee is credited with no completed lifts. 1 - 1 = 0.

Another tool in your quest to capture all the sets and reps you've done is the **+** operator. This symbol allows you to link together movements to record your supersets or rounds.

Blasting your bis and tris? Well, sun's outs, guns out! Let's try it out
```
Bi Blaster 9000: 100 20r 3s
+ Tri Destroyer Supreme: 90 10r 100 10r 110 10r
```
You can chain as many movements together as you feel like doing. The number of sets _should_ match but if you happen to want another set on the Tri Destroyer Supreme, go for it.

For recording rounds, an example could be 20.4 from the Crossfit Open
```
Deadlift: 255 21r 255 15r 255 9r
+ Handstand Push-Up: 21r 15r 9r

Deadlift: 315 21r 315 15r 315 9r
+ Handstand Walk: 50 50 50
```
No matter how metconned you get, **Traindown** can support your recording of it!

## Formatting for Easy Reading

Back to the burrito example...I personally like how succinct it looks but it is difficult to read, especially after three burritos. The same training could be expressed in a more reader-friendly format like
```
@ Dec 25 2019 8:03am

# Week: 2
# Day: 4
# Burritos: 3
# Unit: lbs

Squat:
  135 10r 2s
  225 5r
  315 5r
  405 2r
  455
  500 3s
```
I dunnno about you, but that is _much_ easier on my eyes. Since **Traindown** doesn't specify rules around whitespace, you are free to really do whatever in terms of layout. Just be careful not to assert the superiority of spaces over tabs...

## Notes

So now we have a good idea of what happened on that day, but I'd imagine after eating three burritos, they may have been some...issues along the way. Lucky for us, **Traindown** provides tools for capturing notes. Let's add some!
```
@ Dec 25 2019 8:03am

# Week: 2
# Day: 4
# Burritos: 3
# Unit: lbs

* Definitely should have stopped at two burritos.

Squat:
  135 10r 2s
  225 5r
  315 5r
  405 2r
  455
  500 3s
```
What we just added was a **session note**. This is a note that applies to the entire training session and we know that because it came before any movements, right below the session metadata. These notes are helpful for painting a landscape of your training. Let's add some fine details!
```
@ Dec 25 2019 8:03am

# Week: 2
# Day: 4
# Burritos: 3
# Unit: lbs

* Definitely should have stopped at two burritos.

Squat:
  135 10r 2s
  225 5r
  315 5r
  405 2r
    * Had to stop at two reps to run to the bathroom.
  455
  500 3s
    * Surprisingly easy singles given that episode at 405.
    * Consider going up in weight next week.
```
Remember that each combination of a load for sets and reps is considered a **performance** of a given movement. Notes that appear _after_ a performance are called...you guessed it, **performance notes**. You can apply to whatever granularity of work you like to capture. For example, the three singles at 500 could look something like
```
  500
    * Surprisingly easy single given that episode at 405.
  500
    * Whoops. Spoke too soon. Clenching!
  500
    * Another easy rep.
```
Similar to metadata, there exists a third kind of note available to you called the **movement note**. This note is useful for capturing information about the entire movement such as stance, grip placement, etc. A movement note is declared after the movement name and before any performances, just like the movement metadata. For example
```
Squat:

  * Worked on trying to increase interabdominal pressure, but not too much considering the burritos.
  * Grip was pinky on the rings.

  135 10r 2s
  225 5r
  315 5r
  405 2r
    * Had to stop at two reps to run to the bathroom.
  455
  500 3s
    * Surprisingly easy singles given that episode at 405.
    * Consider going up in weight next week.
```
The notes about pressure and hand placement apply to _all the performances_ of the movement.

## Recap

You now know all there is to know about **Traindown**.

Each session has many performances as well as a date and time on which it occurred. All three types of information-sessions, movements, and performances-can have notes and metadata. The layout of your file doesn't really matter in so far as you can read it years from now or the tool you use to analyze your training can read it.

It's entirely up to you how you want to record your training and I recommend trying out different strategies to see what feels best for you. If you happen to have any questions, please feel free to email me at [tyler at greaterscott dot com](mailto: tyler@greaterscott.com). Kick ass!

## FAQs

### What is the preferred file extension for Traindown?

This will likely depend on your analysis tool of choice, but the standard file extension is **.traindown**—e.g., _20190101.traindown_. You could just as easily save your files as plain text files, if you prefer.

### What if part of my training is done in a different unit?

Ah, a true connoisseur of units, I see. This is well supported using performance metadata and the "Unit" key. For example, lets say we needed to add some kettlebell swings to our session from above.
```
@ Dec 25 2019 8:03am

# Week: 2
# Day: 4
# Burritos: 3
# Unit: lbs

* Definitely should have stopped at two burritos.

Squat:
  135 10r 2s
  225 5r
  315 5r
  405 2r
    * Had to stop at two reps to run to the bathroom.
  455
  500 3s
    * Surprisingly easy singles given that episode at 405.
    * Consider going up in weight next week.

Kettlebell swings: 24 10r 4s
  # Unit: kg
```
### Does Traindown support non-weight based training?

Sadly, we can't lift weights all the time. Sometimes we wind up doing things like cardio or maybe even compete in sports that don't involve weights. **Traindown** is flexible enough to help you record all kinds of training. Let's see an example—this one is for a sprinter.
```
@ 12/25/19

Warmup drills: 3s

60m sprints: 7.12 7.03 7.01
  # Unit: seconds

Jog: 1600
  # Unit: meters
```
Units in **Traindown** can really be _anything_, so whatever it is you are doing, you likely can record it with **Traindown**.

* * *

Made with 🤬. © | 2020 Greater Scott | All rights reserved.

Free for use under the BSD license.
