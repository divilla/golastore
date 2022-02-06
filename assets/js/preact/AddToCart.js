import { h, Component, render } from 'https://unpkg.com/preact?module';

export class AddToCart extends Component {
    render(props, state) {
        return h('p', { 'class': 'control' }, [
            h('button', { 'class': 'button is-small' }, [
                h('span', { 'class': 'icon is small' }, [
                    h('i', { 'class': 'fas fa-minus' })
                ])
            ])
        ])
    }
}
