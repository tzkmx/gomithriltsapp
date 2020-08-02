import m from "mithril";
import { withKnobs } from "@storybook/addon-knobs";
import { withA11y } from "@storybook/addon-a11y";
import { Menu } from "./menu";

export default {
  title: "Component/Menu",
  component: Menu,
  decorators: [withKnobs, withA11y],
};

export const menu = (): m.Component => ({
  view: () => m(Menu),
});
