var root = document.body;
m.render(root,  [
    m("main", [
        m("h1", {class: "title"}, "My first app"),
        m("a", {class: "button"}, "ボタン"),
        m("i", {class: "fas fa-jedi"}),
        m("i", {class: "fab fa-github"})
    ])
]);
