import './scss/theme.scss'
import octicons from 'octicons'
import Bookmark from './bookmark'

// Generate octicons
let octiconsEl = document.querySelectorAll('.octicon')
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
  })
})

document.querySelector('.js-b-indicator').addEventListener('click', () => {
  let bookmarkNode = document.querySelector('.bookmark')
  if (!bookmarkNode.classList.contains('show')) {
    bookmarkNode.classList.add('show')
  } else {
    bookmarkNode.classList.remove('show')
  }
})

document.querySelectorAll('.js-b-additem').forEach(node => {
  node.addEventListener('click', (e) => {
    e.preventDefault()
    let items = b.data.items
    let link = node.getAttribute('data-link')
    if (items.filter(item => item.link === link).length > 0) {
      return
    }
    node.classList.add('added')

    items.push({
      link: link,
      title: node.getAttribute('data-title'),
      date: node.getAttribute('data-date'),
      site: node.getAttribute('data-site')
    })

    b.data.items = items

    localStorage.setItem('bookmark_items', JSON.stringify(items))
  })
})
