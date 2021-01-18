---
title: Spec
template: default.html
name: spec
---

<h1 class="hero center-text">Spec</h1>
<h2 class="center-text">The Traindown language specification.</h2>
<h2 class="center-text">Version 1.2</h2>
<small class="block center-text">Published 2020.01.18</small>

---

This specification is intended for **developers wanting to build software that uses Traindown**.

*If you are a consumer looking for information on how to use **Traindown** to record your training, please refer to the [Guide](/guide).*

## Overview

The specification presented below aims to provide a complete definition of the **Traindown** Language such that you may go out into the world and develop software that can help other folks record and make sense of their personal training history.

The language used in this specification is informal but comprehensive. Should you spot any errors or have any questions, please don't hesitate to email me at [tyler at greaterscott.com](mailto:tyler@greaterscott.com).

Also, don't forget to check out the [Source](/source) page and the repos linked there for official prebuilt libraries that make using **Traindown** in your project as easy as adding another dependency.

Without further ado, here is **Traindown**.

## Keywords and Identifiers

**Traindown** has a relatively small set of reserved keywords and identifiers.

There is intention behind this. The point is to require as little cognitive overhead of the person (not user!) as possible. You have to figure that alot of **Traindown** is typed while out of breath and covered in chalk or whatever else.

<table>
  <thead><tr><th>Token</th><th>Name</th><th>Rules</th></tr></thead>
  <tbody>
    <tr>
      <td>EOF</td>
      <td>End of file</td>
      <td><pre>System defined</pre></td>
    </tr>
    <tr>
      <td>LT</td>
      <td>Line Terminator</td>
      <td><pre>";" | Unicode line break codepoints</pre></td>
    </tr>
    <tr>
      <td>ANY</td>
      <td>Non Line Terminating Unicode</td>
      <td><pre>!LT</pre></td>
    </tr>
    <tr>
      <td>NUM</td>
      <td>Number Values</td>
      <td><pre>0123456789.</pre></td>
    </tr>
    <tr>
      <td>DT</td>
      <td>Date / Time</td>
      <td><pre>"@" ANY* LT</pre></td>
    </tr>
    <tr>
      <td>MK</td>
      <td>Meta Key</td>
      <td><pre>"#" ANY* ":"</pre></td>
    </tr>
    <tr>
      <td>MV</td>
      <td>Meta Value</td>
      <td><pre>MK ANY* LT</pre></td>
    </tr>
    <tr>
      <td>NOTE</td>
      <td>Note</td>
      <td><pre>"*" ANY* LT</pre></td>
    </tr>
    <tr>
      <td>MN</td>
      <td>Movement Name</td>
      <td><pre>ANY* :</pre></td>
    </tr>
    <tr>
      <td>SSMN</td>
      <td>Superset Movement Name</td>
      <td><pre>"+" ANY* :</pre></td>
    </tr>
    <tr>
      <td>FAILS</td>
      <td>Failures</td>
      <td><pre>NUM* ("f" | "F")</pre></td>
    </tr>
    <tr>
      <td>REPS</td>
      <td>Reps</td>
      <td><pre>NUM* ("r" | "R")</pre></td>
    </tr>
    <tr>
      <td>SETS</td>
      <td>Sets</td>
      <td><pre>NUM* ("s" | "S")</pre></td>
    </tr>
    <tr>
      <td>LOAD</td>
      <td>Load</td>
      <td><pre>NUM* | ("b" | "B")("w" | "W") ["+" NUM*]?</pre></td>
    </tr>
  </tbody>
</table>

**Traindown** does not have composition of tokens into compound types or values. The table above represnts everything you will need to tokenize in order to parse a **Traindown** document

## Scopes

There are **four** scopes to a **Traindown** document ordered here by broadest to narrowest:

> **Document**
> *From the beginning of the file to the **EOF** token.*

> **Session**
> *From a **DT** token until **EOF** or another **DT** token*

> **Movement**
> *From an **MN/SSMN** token until **EOF** or another **MN/SSMN** token*

> **Performance**
> *From a **LOAD** token until **EOF** or another **LOAD** token.*

### Document Scope

Currently this scope is not used in the parsing of **Traindown**.

Should your software make use of the document scope, please be sure to provide the people using your software of this. Also be aware that if you make use of the scope while it is currently undefined, you may have incompatibility with a future version of **Traindown**.

That said, if you are using it, please reach out an let me know so that we can work to bring it into the spec.

### Session Scope

A **session** is to be considered as a single training event.

Every session **must** have at least a date on which it occurred.

You **should** provide error messaging to remind the person that they should include a date (and/or time) in their **Traindown** document.

A single **Traindown** document **may** contain more than one session. This is up to the person writing the document. Should your software not honor this, please be sure to message this to the people using your software. Ideally this is to be done prior to an error message and you should offer to auto format into one session per document as this is simple to do.

In addition to the required date (and/or time), a session **may** contain metadata key/value pairs, notes, and movements. All three of those should be considered *optional*. A session without movements may not seem like a use case, but it is.

### Movement Scope

A **movement** is to be considered as a single exercise or event in a training session. An example of this would be something like "squats".

A movement **may** contain metadata key/value pairs, notes, and performances.

Like with sessions and movements, it may not make sense to have a movement without performances, but it is a supported use case so let's honor it.

### Performance Scope

A **performance** is the expression of a movement.

Each performance **must** have a load.

A person writing **Traindown** may optionally provide sets, reps, and failures. In absence of explicit detail, you should assume a single set, a single rep, and no failures.

A performance **may** contain metadata key/value pairs and notes.
