import './scss/theme.scss'
import octicons from 'octicons'
import Bookmark from './bookmark'

const toggleDropdown = () => {
  let bookmarkNode = document.querySelector('.bookmark')
  if (!bookmarkNode.classList.contains('show')) {
    bookmarkNode.classList.add('show')
  } else {
    bookmarkNode.classList.remove('show')
  }
}

// Generate octicons
let octiconsEl = document.querySelectorAll('.i-octicon')
octiconsEl.forEach((oct) => {
  let icon = oct.getAttribute('data-icon')
  if (octicons[icon]) {
    oct.innerHTML = octicons[icon].toSVG()
  }
})

// Initialize the bookmark
let b = Bookmark({
  items: localStorage.getItem('bookmark_items') ? JSON.parse(localStorage.getItem('bookmark_items')) : []
})

// Highlight the bookmark icon of added items.
b.data.items.forEach(item => {
  document.querySelectorAll(`[data-link="${item.link}"]`).forEach(node => {
    node.classList.add('added')
    node.textContent = 'remove bookmark'
  })
})

document.querySelector('.js-b-indicator').addEventListener('click', () => {
  toggleDropdown()
})

// Add/remove bookmark
document.querySelectorAll('.js-b-additem').forEach(node => {
  node.addEventListener('click', (e) => {
    e.preventDefault()
    let items = b.data.items
    let link = node.getAttribute('data-link')
    // Add item
    if (!node.classList.contains('added')) {
      if (items.filter(item => item.link === link).length > 0) {
        return
      }

      node.classList.add('added')
      node.textContent = 'remove bookmark'

      items.push({
        link: link,
        title: node.getAttribute('data-title'),
        date: node.getAttribute('data-date'),
        site: node.getAttribute('data-site')
      })

      b.data.items = items
    } else { // Remove item
      items = b.data.items.filter(item => {
        return item.link !== link
      })

      node.classList.remove('added')
      node.textContent = 'bookmark'

      b.data.items = items
    }

    localStorage.setItem('bookmark_items', JSON.stringify(items))
  })
})

// Remove bookmark in the dropdown list.
document.addEventListener('click', (e) => {
  e.preventDefault()

  if ( ! e.target.classList.contains('js-b-dropdown-remove-bookmark') ) {
    return
  }

  let items = b.data.items.filter(item => {
    return item.link !== e.target.getAttribute('data-id')
  })

  b.data.items = items

  localStorage.setItem('bookmark_items', JSON.stringify(items))
})
