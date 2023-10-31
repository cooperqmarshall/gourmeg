function handelListSearchResultClick(e) {
  let list = e.srcElement.innerText
  document.getElementById('list-input').value = list
}

// const listInputElement = document.getElementById('list-input')
// listInputElement.addEventListener('focus', () => {
//   const listDropdownElement = document.getElementById('list-dropdown')
//   listDropdownElement.classList.add('active')
//   listDropdownElement.classList.remove('inactive')
// })


