const Bookmark = (dataObj) => {
  observeData(dataObj)

  return {
    data: dataObj
  }

  function makeReactive (obj, key) {
    let val = obj[key]

    Object.defineProperty(obj, key, {
      get () {
        return val
      },
      set (newVal) {
        val = newVal
        parseDOM(obj)
      }
    })
  }

  function observeData (obj) {
    for (let key in obj) {
      if (obj.hasOwnProperty(key)) {
        makeReactive(obj, key)
      }
    }

    parseDOM(obj)
  }

  function parseDOM (obj) {
    let count = obj.items.length
    let dropdownNode = document.querySelector('.js-b-dropdown')

    // update counter
    document.querySelector('.js-b-counter').textContent = count

    // render bookmark items
    if (count === 0) {
      dropdownNode.innerHTML = '<span class="no-item">No bookmark items found.</span>'
    } else {
      let html = ''
      obj.items.forEach(item => {
        html += `
        <a href="${item.link}" class="link" target="_blank" rel="nofollow noopener">
          <h1>${item.title}</h1>
          <div class="meta">${item.date} by ${item.site}</div>
        </a>
        `
      });
      dropdownNode.innerHTML = html
    }
  }
}

module.exports = Bookmark
