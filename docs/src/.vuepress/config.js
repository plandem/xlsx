const mdList = require('markdown-it-task-lists');

module.exports = {
    title: 'Xlsx2Go',
    description: 'Fast and Reliable way to work with xlsx in Golang',
    markdown: {
        extendMarkdown: md => md.use(mdList)
    },
    themeConfig: {
        logo: '/logo.png',
        repo: 'plandem/xlsx',
        repoLabel: 'GitHub',
        editLinks: true,
        editLinkText: 'Help to improve this page!',
        lastUpdated: 'Last Updated',
        displayAllHeaders: true,
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
                    '/guide/merged-cells.md',
                    '/guide/hyperlinks.md',
                    '/guide/comments.md',
                    '/guide/styles-formatting.md',
                    '/guide/conditional-formatting'
                ]
            }
        ]
    }
};
