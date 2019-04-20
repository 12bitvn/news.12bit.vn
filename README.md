# news.12bit.vn

Browse multiple Vietnamese tech blogs at the same time.

[![Netlify Status](https://api.netlify.com/api/v1/badges/51d69537-8760-4eef-ad70-fead35768f13/deploy-status)](https://app.netlify.com/sites/jovial-spence-90afc6/deploys)

![featured-image](https://user-images.githubusercontent.com/3280351/56430895-1a0a4880-62f2-11e9-98b2-d3bb66ae80ab.png)

## Contributing

**Prerequisites**

Since Netlify does not support Hugo extended version that allows Hugo converts SCSS to CSS. We need to use a task runner to convert SCSS. In this case, webpack is our first choice.

- Hugo
- Node.js, npm

**Install Node.js dependencies**

```
yarn install
```

**Start webpack development & local server**

You may need to run the both commands at the same time.

```
yarn dev
hugo serve
```

Browse the site at http://localhost:1313

## Add your site

> Trust me, it's really simple. You don't need to inbox us.

If you have a blog, you can add your RSS link into [this file](https://github.com/12bitvn/news.12bit.vn/blob/master/data/links.json), then create a pull request. We'll review it as soon as possible.
