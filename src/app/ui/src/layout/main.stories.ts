import m from "mithril";
import { withKnobs, text } from "@storybook/addon-knobs";
import { withA11y } from "@storybook/addon-a11y";
import LayoutMain from "./main";
import SimplePage from "@/component/simple-page";

export default {
  title: "Component/Layout Main",
  component: LayoutMain,
  decorators: [withKnobs, withA11y],
};

export const simplePage = (): m.Component => ({
  view: () => {
    return m(
      LayoutMain,
      m(SimplePage, {
        title: text("Title", "This is the Title"),
        description: m(
          "span",
          text("Description", "This is a subtitle or description.")
        ),
        content: text("Content", "This is the content."),
      })
    );
  },
});