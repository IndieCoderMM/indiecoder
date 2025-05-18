// Disable scrolling when the menu is open
const menuToggle = document.getElementById('menu-toggle');
const body = document.body;
menuToggle.addEventListener('change', function() {
  if (menuToggle.checked) {
    body.style.overflow = 'hidden';
  } else {
    body.style.overflow = '';
  }
});

console.log("hello main")
