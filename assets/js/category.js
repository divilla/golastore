import { h, Component, render } from 'https://unpkg.com/preact?module';
import { AddToCart } from './preact/AddToCart.js'

// const app = h(AddToCart, null, 'Title');

render(h(AddToCart), document.getElementsByName('add-to-cart'));
// render(link, document.body);
