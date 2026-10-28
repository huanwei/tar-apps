# Completion Recovery Drill

This note documents the completion failure drill for the TAR pilot service.

## Goal

Verify that the control plane can recover a work item when a worker dies
mid-completion, without losing work or leaving the workspace in a broken state.

## Procedure

1. Start a work item from the base snapshot and make a change in the workspace.
2. Stop the worker from renewing its lease so the attempt expires.
3. Let the control plane expire the lease and recover the attempt.
4. Re-run the work item from the base snapshot with a fresh worker.
5. Run the acceptance checks for the item to confirm the recovered workspace
   is runnable.

## Expected result

The expired attempt is recovered, the work item is re-run from the base
snapshot, and the acceptance checks pass on the recovered workspace. All work
happens inside this workspace; there is no Git metadata to clean up.
