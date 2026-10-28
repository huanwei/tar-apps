# Budget drill

This file is the artifact of a budget-limit drill: it is created while the
per-work-item budget is set very low.

## Fail-closed behavior

When the budget for a work item is exhausted, the system fails **closed**, not
open. That means:

- Further work on the item is halted rather than allowed to continue
  unmeasured.
- No partial or unverified changes are promoted as if they had succeeded.
- The item is surfaced for a human decision instead of silently proceeding.

Failing closed keeps an over-budget run from producing unbounded side effects
or unreviewed results. The safe default under an exceeded budget is to stop and
report, not to push forward.
