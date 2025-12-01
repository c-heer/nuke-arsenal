// Simple reactive store for triggering refreshes
let refreshCounter = $state(0);

export const refreshTrigger = {
  get value() { return refreshCounter; },
  trigger() { refreshCounter++; }
};

export function triggerRefresh() {
  refreshCounter++;
}
