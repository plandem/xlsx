const mdList = require('markdown-it-task-lists');
const path = require('path');

module.exports = {
    base: '/xlsx2go/',
    title: 'Xlsx2Go',
    description: 'Fast and Reliable way to work with xlsx in Golang',
    themeConfig: {
        logo: '/logo.png',
        repo: 'plandem/xlsx',
        repoLabel: 'GitHub',
        editLinks: false,
        // editLinkText: 'Help to improve this page!',
        lastUpdated: 'Last Updated',
        displayAllHeaders: true,
        algolia: {
            apiKey: '3e97b85acaa0479ed415ab8ecdaf55d6',
            indexName: 'xlsx2go'
        },
        nav: [
            {text: 'Home', link: '/'},
            {text: 'Guide', link: '/guide/'},
            {text: 'API', link: 'https://godoc.org/github.com/plandem/xlsx'}
        ],
        sidebar: [
            {
                title: 'Guide',
                collapsed: false,
                children: [
                    '/guide/',
                    '/guide/getting-started.md',
                    '/guide/access.md',
                    '/guide/iterate.md',
                    '/guide/values',
                    '/guide/merged-cells.md',
                    '/guide/rich-text.md',
                    '/guide/hyperlinks.md',
                    '/guide/comments.md',
                    '/guide/styles-formatting.md',
                    '/guide/conditional-formatting',
                    '/guide/copy',
                    '/guide/streams.md',
                    '/guide/filters.md',
                    '/guide/limits.md',
                ]
            },
            {
                title: 'F.A.Q.',
                path: '/faq/',
            }
        ]
    },
    markdown: {
        extendMarkdown: md => md.use(mdList)
    },
    chainWebpack: (config, isServer) => {
        config.resolve.alias.set('@images', path.resolve(__dirname, "./public"))
    },
    plugins: [
        '@vuepress/back-to-top',
        '@vuepress/last-updated',
        '@vuepress/nprogress',
        ['container', {type: 'right', defaultTitle: ''}],
        ['container', {type: 'note', defaultTitle: ''}],
        ['@vuepress/google-analytics', {'ga': 'UA-122513-16'}],
    ],
};
