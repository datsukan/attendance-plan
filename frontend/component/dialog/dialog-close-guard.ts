let lastCloseTime = 0;
const GUARD_MS = 150;

export const markDialogClosed = () => {
  lastCloseTime = Date.now();
};

export const isDialogRecentlyClosed = () => {
  return Date.now() - lastCloseTime < GUARD_MS;
};
