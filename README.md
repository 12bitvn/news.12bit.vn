# news.12bit.vn

Browse multiple Vietnamese tech blogs at the same time.

[![Netlify Status](https://api.netlify.com/api/v1/badges/51d69537-8760-4eef-ad70-fead35768f13/deploy-status)](https://app.netlify.com/sites/jovial-spence-90afc6/deploys)

## Contribute

**Prerequisites**

Since Netlify does not support Hugo extended version that allows Hugo converts SCSS to CSS. We need to use a task runner to convert SCSS. In this case, webpack is our first choice.

- Hugo
- Nodejs, npm

**Install Node.js dependencies**

```
yarn install
```

**Start the local server**

```
yarn dev
hugo serve
```

Browse the site at http://localhost:1313

## Add your site

> Trust me, it's really simple. You don't need to inbox us.

If you have a blog, you can add your RSS link into [this file](https://github.com/12bitvn/news.12bit.vn/blob/master/data/links.json), then create a pull request. We'll review it as soon as possible.
