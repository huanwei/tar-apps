# Runner recovery drill

Purpose: confirm that losing a runner mid-attempt does not lose the work item.

## Drill

1. Start a work item on a runner and confirm it holds the lease.
2. Kill the runner process (simulating a crashed or partitioned worker).
3. Do not renew the lease; let it expire.
4. Observe the control plane reclaim the item and dispatch it to another worker.
5. Verify the replacement worker completes the item from the base snapshot.

## Pass criteria

- The killed runner's lease expires and is not silently extended.
- Another worker picks up the item without operator intervention.
- The item finishes on the replacement worker and the result is committed.

## Notes

An attempt recovered this way starts from the base snapshot, so any partial
changes made by the killed runner are discarded and must be redone.
