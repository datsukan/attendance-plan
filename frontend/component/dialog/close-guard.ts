let lastCloseTime = 0;
const GUARD_MS = 150;

export const markDialogClosed = () => {
  lastCloseTime = Date.now();
};

export const isDialogRecentlyClosed = () => {
  return Date.now() - lastCloseTime < GUARD_MS;
};

let lastPopoverCloseTime = 0;
const POPOVER_GUARD_MS = 600;

export const markPopoverClosed = () => {
  lastPopoverCloseTime = Date.now();
};

export const isPopoverRecentlyClosed = () => {
  return Date.now() - lastPopoverCloseTime < POPOVER_GUARD_MS;
};
