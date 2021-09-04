/**
 * Checks if the 2nd element is children of 1st parent element
 */
export const isDescendant = (parent, child) => {
  if (!child) return false;
  let node = child.parentNode;
  while (node) {
    if (node === parent) {
      return true;
    }

    // Traverse up to the parent
    node = node.parentNode;
  }

  // Go up until the root but couldn't find the `parent`
  return false;
};

// These should be ideally moved to define plugin after https://github.com/razorpay/frontend-universe/issues/493
export const isClient = () => typeof window === "object";
export const isServer = () => !isClient();
