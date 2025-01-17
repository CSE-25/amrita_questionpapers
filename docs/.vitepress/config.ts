import { defineConfig } from 'vitepress';

// refer https://vitepress.dev/reference/site-config for details
export default defineConfig({
  base: '/amrita_pyq/',
  lang: 'en-US',
  title: 'Amrita PYQ',
  description: 'A CLI tool for accessing PYQs.',
  cleanUrls: true,

  themeConfig: {
    search: {
      provider: 'local',
      placeholder: 'Search',
    },
    nav: [
      {
        text: 'Home',
        link: '/',
      },
      {
        text: 'Docs',
        link: '/getting-started',
      },
    ],
    socialLinks: [
      {
        icon: 'github',
        link: 'https://github.com/CSE-25/amrita_pyq',
      },
    ],

    sidebar: [
      {
        items: [
          { text: 'Gettting started', link: '/getting-started' },
          // { text: 'Examples', link: '/examples' },
        ],
      },
    ],
  },
});
