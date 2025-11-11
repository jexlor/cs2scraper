// that's the main script that should be injected and return data from the page.
// you can modify selectors here if you use different provider (or provided site changes classes).

Array.from(document.querySelectorAll('.item-box')).map(container => ({
weapon: container.querySelector('.collection-skinbox span')?.textContent.trim() || '',
name: container.querySelector('.block:not(.txt-small):not(.txt-stattrak)')?.textContent.trim() || '',
rarity: container.querySelector('.classified, .covert, .restricted, .mil-spec')?.textContent.trim() || '',
price: container.querySelector('.block.txt-small:not(.txt-stattrak):not(.txt-dark-grey)')?.textContent.trim() || '',
stattrakPrice: container.querySelector('.block.txt-small.txt-stattrak')?.textContent.trim() || '',
collection: container.querySelector('.block:not(.txt-small):not(.txt-stattrak):not(.txt-white)')?.textContent.trim() || '',
url: container.querySelector('img')?.src?.replace('/webp/', '/').replace('.webp', '.png') || '' }))