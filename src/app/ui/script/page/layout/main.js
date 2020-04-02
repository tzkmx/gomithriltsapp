import m from "mithril";
import Menu from "../component/menu";

var View = {
  view: function (vnode) {
    return m("main.layout", [m(Menu), m("section", vnode.children)]);
  },
};

export default View;
