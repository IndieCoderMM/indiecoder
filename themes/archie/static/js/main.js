// Disable scrolling when the menu is open
const menuToggle = document.getElementById('menu-toggle');
const body = document.body;
menuToggle.addEventListener('change', function () {
  if (menuToggle.checked) {
    body.style.overflow = 'hidden';
  } else {
    body.style.overflow = '';
  }
});

// Scroll to top button
const topBtn = document.getElementById('topBtn');

// Show the button when the user scrolls down 20px from the top
window.onscroll = function () {
  if (document.body.scrollTop > 20 || document.documentElement.scrollTop > 20) {
    topBtn.style.display = "block";
  } else {
    topBtn.style.display = "none";
  }
}

// When the user clicks on the button, scroll to the top of the document
topBtn.addEventListener('click', function () {
  window.scrollTo({ top: 0, behavior: 'smooth' });
});


