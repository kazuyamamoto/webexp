var root = document.body;
m.render(root, [
    m("section", {class: "section"},
        m("div", {class: "container"}, [
            m("h1", {class: "title"}, "Hello world"),
            m("p", {class: "subtitle"}, [
                "My first website with ",
                m("strong", "Bulma"),
                "!"
            ]),
            m("p", m("a", {class: "button"}, "ボタン")),
            m("p", [
                m("i", {class: "fas fa-jedi"}),
                m("i", {class: "fab fa-github"})
            ])
        ])
    )
]);
