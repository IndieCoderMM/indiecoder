---
title: My Rubber Duck Has Read the Docs
description: How I fixed a weird React bug using ChatGPT
date: 2025-06-07
tags: [ai, productivity, react]
---

Today while I was building a number memory game with React, I encountered a strange behaviour with input focus. 

There was a grid of input fields and I wanted to auto-focus next inputs sequentially after entering a number - a standard OTP-input feature. The first input was auto-focused on mount with `useEffect` so there was no problem. But the auto-focus just wouldn't work on the second input. However, once the second one was focused manually, the rest of the flow worked fine.

At first, I thought this might be a `type="number"` issue so I changed it and it didn't work. Another strange thing was that this issue wasn't specific to the second input. No matter where I started, say 50th input, the very next one always needed manual focus.

At the point, I did what any real 10x developer would do. I opened up *ChatGPT* and started questioning its ability. Jokes aside, it's basically my go-to brower for searching anything. I first asked for a general solution, it gave me a few suggestions, but none solved it. So, I described my setup and explain my use case with code snippets. After about 5 rounds of back and forth (about an hour), we found a solution.

The fix was to wrap the focus in `requestAnimationFrame` like this. See [complete code here](https://github.com/IndieCoderMM/memio/commit/eb32cb36d58915f16f26138c93ab61985a3c918d#diff-57189158dcaf4f173b60c8743b60097ee78d3075fef44a3ae080a633d738d9f4R25).

```js
requestAnimationFrame(() => {
  nextInput.focus();
});
```

It works similar to `setTimeout` and defers a function call untill:
- The current paint cycle completes.
- All DOM updates are committed.
- Browser is done handling layout, keyboard etc.

I tried using `setTimeout`, but it didn't work. And, the issue only happening on the first focus is likey due to browser's internal focus logic and a cold start case, but I'm not planning to do research on this. I'm not even sure if this is a correct approach, but I don't mind as long as it solves my problem.

What I really took away from this is *how AI will replace us*. I believe **AI is going to replace developers, but *only in boring parts**. This kind of oddly specific bug could take hours to figure out on my own, which would involve a lot of reading, diving into StackOverflow or getting lost in endless threads on GitHub. Within a few seconds, *ChatGPT* could do all the research and come up with seemingly relevant solutions.

Of course, dumping a buggy codebase into a prompt and yelling at an LLM to *just fix it* won't work. Instead, I treat it like  *a rubber duck*. You've probably heard of **rubber duck debugging**. Now, imagine what you could do if your cute little duck could talk back, and had access to the entire knowledge base of humankind. You explain your problem, it responds with suggestions, edge cases, or even ideas you hadn’t considered. That's how I use AI, **not to replace thinking, but to boost it**.

It won’t always get things right. But even when it’s wrong, it sparks the right questions. And honestly, that’s sometimes all you need to get unstuck.

