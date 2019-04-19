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
      dropdownNode.innerHTML = `
      <p class="no-item">
        <span>No bookmark items found.</span>
      </p>
      `
    } else {
      let html = ''
      obj.items.forEach(item => {
        html += `
        <div class="link" >
          <h1><a href="${item.link}" target="_blank" rel="nofollow noopener">${item.title}</h1>
          <div class="meta">
            <a href="http://${item.site}" target="_blank" rel="nofollow noopener">${item.site}</a> | <a href="#" class="js-b-dropdown-remove-bookmark" data-id="${item.link}">remove</a>
          </div>
        </div>
        `
      });
      dropdownNode.innerHTML = html
    }
  }
}

module.exports = Bookmark
