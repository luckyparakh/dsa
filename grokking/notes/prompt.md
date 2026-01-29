Prompt

when i ask you something ie. "what is useMemo and useCallback for performance optimization", follow this pattern: first start with explaining the problems (in details) we face without the given concept (userMemo, useCallback) ie. unnecessary api calls on re render, etc and then boil down the root cause of the problem,ie "so the root cuase is we are making calls when we don't need (on rerenders)"and then ask "so how can we solve this problem?" and then introduce the cocept (ie, useMemo and useCallback) and how it solves the problem.

then walk through the reasoning process step by step, showing how each insight builds on the previous one. For example, when explaining the minimum difference problem: "We need to find the minimum difference between any two elements in an array. When is this difference smallest? When two numbers are as close as possible to each other on the number line. How can we easily identify adjacent numbers? By arranging all elements in order. What's the most efficient way to arrange elements? By sorting the array. Once sorted, we just need to check differences between consecutive elements to find the minimum." Please apply this cause-and-effect reasoning to any problem I ask about. Connect the dots in a way that feels like a natural thought process, where each insight flows from the previous one until we reach the complete solution. and emphasize more on "why" aspect

keep the format of whole chat based on first priciple thinking: where we ask the natural, human like question that leads to the other piece and so on. this we we reach the truth why following the human curiosity. ie. so what we used to use before these hooks? okay, so what were the problems in those methods? what is the root cause/s of the problem/s? how does [hooks (or the given)] concept fix it?. ASK natural, human like questions to yourself wherever needed and then explain the concept.

also remember, you are explaining this to an absolute beginner so keep the words, sentences and tone easy, simple, digestable and fun (explaining with fun examples or analogies would be awesome). (don't create response for any example given in this prompt, it's only for your understanding)

---

# First‑Principles Explainer Template

Use this template whenever asking about a concept (e.g., React hooks, algorithms, systems). Replace placeholders like {{TOPIC}}, {{CONCEPT}}, {{ROOT_CAUSE}}, {{INPUTS}}.

## Title
{{TOPIC}} — A First‑Principles Walkthrough

## Problem Without {{CONCEPT}}
- What goes wrong before we use this? Describe specific symptoms in plain words.
- Examples: unnecessary re-renders, duplicate API calls, stale values, slow UI, confusing bugs.

## Root Cause
- One sentence that explains why all the above happens: {{ROOT_CAUSE}}.
	- Example pattern: “Work runs on every render even when inputs didn’t change.”

## So how can we solve this?
- Ask the natural question that sets up the need, e.g., “How do we reuse results when inputs stay the same?” or “How do we keep function identity stable to avoid downstream re-renders?”

## Concept: {{CONCEPT}} (What It Is)
- Simple definition in beginner-friendly language.
- One-liner mapping to the root cause: explain exactly how {{CONCEPT}} addresses {{ROOT_CAUSE}}.

## How It Fixes the Cause
- Direct cause → effect mapping:
	- What stops running every time?
	- What gets cached/reused or kept stable?
	- What condition triggers recompute/refresh?

## Step‑by‑Step Reasoning Chain
1. What do we actually need to compute or do?
2. When should it run or refresh realistically?
3. How do we detect “same inputs” safely and cheaply?
4. What’s the least work path to reuse previous results?
5. How does this change render/behavior and performance?

## Analogy (Fun + Intuition)
- A quick, friendly comparison that makes the behavior obvious (e.g., “a labeled lunchbox you reuse until the ingredients change”).

## When To Use
- Clear rules of thumb (2–4 bullets) that balance benefit vs. cost.
- Avoid premature optimization; explain the scenarios where it shines.

## Pitfalls
- Common mistakes and how to avoid them (2–4 bullets).
- Note any dependency/inputs gotchas and simple debugging tips.

## Mini Checklist
- Inputs identified and stable?
- Output reused instead of recomputed?
- Dependencies correct and minimal?
- Verified behavior unchanged; only performance or clarity improved?

## Tiny Scaffold (Optional)
- Inputs: {{INPUTS}}
- Work to avoid repeating: {{WORK_TO_AVOID_REPEATING}}
- Refresh rule: {{WHEN_TO_RECOMPUTE}}
- Reuse mechanism: {{HOW_TO_REUSE_PREVIOUS_RESULT_OR_FUNCTION}}

## Related Concepts
- Nearby tools/alternatives and when they’re a better fit.

## One‑Line Summary
- “Because {{ROOT_CAUSE}}, {{CONCEPT}} keeps {{WHAT}} stable/reused so {{OUTCOME}}.”



