---
title: Custom Theme Switcher Hook in React
description: Learn how to build a reusable theme switcher hook in React to seamlessly toggle between dark and light modes, with persistence across sessions.
date: 2023-10-04
tags: [react, frontend, webdev, js]
---

In this tutorial, I’ll show you how to build a custom theme switcher hook in React. By the end, you’ll have a reusable hook that can add dark/light mode to any React project.

We’ll cover:  
1. Creating the hook step-by-step.  
2. Saving the user’s theme preference.  
3. Adding dark/light mode toggle functionality.  

## Logic

Here’s how the custom hook works:

1. It checks for a saved theme in local storage.  
2. If no saved theme is found, it checks the system’s dark mode preference.  
3. It updates the theme dynamically and stores the user’s choice for the next session.  

## Creating a custom hook

This custom hook will handle the logic for switching between light and dark themes. Let’s create a new file called `useThemeSwitcher.js` and add the following code:

```js
import { useEffect, useState } from "react";

const useThemeSwitcher = () => {
  // Define the media query for dark mode preference
  const preferDarkQuery = "(prefer-color-scheme: dark)";

  // Initialize state to store the current theme
  const [mode, setMode] = useState("");

  // Use the useEffect hook to set the initial theme
  useEffect(() => {
    // Check if a theme is stored in local storage
    const localMode = localStorage.getItem("theme");

    // Check if the user prefers dark mode
    const prefersDark = window.matchMedia(preferDarkQuery).matches;

    // Set the initial theme based on local storage or user preference
    if (localMode) {
      setMode(localMode);
    } else if (prefersDark) {
      setMode("dark");
    } else {
      setMode("light");
    }
  }, []);

  // Use another useEffect hook to update the theme when it changes
  useEffect(() => {
    // Update local storage with the current theme
    if (mode === "dark") {
      localStorage.setItem("theme", "dark");
      document.documentElement.classList.add("dark");
    } else {
      localStorage.setItem("theme", "light");
      document.documentElement.classList.remove("dark");
    }
  }, [mode]);

  // Return the current theme and a function to switch themes
  return [mode, setMode];
};

export default useThemeSwitcher;
```

- `preferDarkQuery` variable is set to check if the user prefers dark mode. It uses a special CSS media query.
- We initialize the `mode` state variable to keep track of the current theme (light or dark).
- The first `useEffect` checks local storage and the user’s system preferences to set the initial theme on app load.
- The second `useEffect` will run whenever we change the theme using `setMode`, and it will update the app’s theme and saves the user’s choice to local storage.

## Using the Hook

You can use this hook in your React components like this:

```js
import useThemeSwitcher from "./useThemeSwitcher";

function App() {
  const [theme, setTheme] = useThemeSwitcher();

  // Create a function to toggle the theme
  const toggleTheme = () => {
    if (theme === "light") {
      setTheme("dark");
    } else {
      setTheme("light");
    }
  };

  return (
    <div className={`App ${theme}`}>
      <h1>Theme Switcher Example</h1>
      <button onClick={toggleTheme}>Toggle Theme</button>
      <p>Current Theme: {theme}</p>
    </div>
  );
}

export default App;
```

And that’s it! You now have a reusable custom hook for managing dark/light mode in your React apps.

Happy coding! 
