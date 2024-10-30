function openTab(evt, tabName) {
    var i, tabContent, tabLinks;
    tabContent = document.getElementsByClassName("tab-content");
    for (i = 0; i < tabContent.length; i++) {
        tabContent[i].style.display = "none";
    }
    tabLinks = document.getElementsByClassName("tab");
    for (i = 0; i < tabLinks.length; i++) {
        tabLinks[i].className = tabLinks[i].className.replace(" active", "");
    }
    document.getElementById(tabName).style.display = "block";
    evt.currentTarget.className += " active";
}

const cursor = document.createElement('div');
cursor.classList.add('custom-cursor');
document.body.appendChild(cursor);

document.addEventListener('mousemove', (e) => {
  cursor.style.left = e.clientX + 'px';
  cursor.style.top = e.clientY + 'px';
  
  // Check if the element under cursor is clickable
  const elementUnderCursor = document.elementFromPoint(e.clientX, e.clientY);
  if (elementUnderCursor) {
    const isClickable = (
      elementUnderCursor.matches('a, button, [role="button"], input, .clickable') ||
      elementUnderCursor.closest('a, button, [role="button"], input, .clickable')
    );
    
    // Set opacity based on whether element is clickable
    cursor.style.opacity = isClickable ? '0' : '1';
  }
});