import { h, Component, render } from 'https://unpkg.com/preact?module';

export class Link extends Component {
    render(props, state) {
        return h('a', {href: props.href}, ...props.children)
    }
}
