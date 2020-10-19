<template>
	<div :class="`card-list__two-col card-list card-list__two-col--${$route.params.uid}`">
		
		<div class="card-list__two-col--col">
			<div class="card-item" v-for="(card, i) in oddCards" :key="i" @click="$set(card, 'selected', !card.selected)" :class="{ 'card-item--active': card.selected }">
				<v-card class="card-item__card" elevation="10" light>
					<font-awesome-layers class="fa-3x card-item__icon" v-if="card.icon">
						<font-awesome-icon class="card-item__icon--background" :icon="['fas', 'circle']" />
						<font-awesome-icon class="card-item__icon--plus" :icon="['fas', 'plus']" transform="shrink-6" />
					</font-awesome-layers>
					<v-img class="card-item__background" dark :src="card.card_bg.url" :alt="card.card_bg.alt">
						<h2 class="card-item__title">{{ $prismic.asText(card.card_title) }}</h2>
						<prismic-rich-text :field="card.card_text" class="card-item__text" />
					</v-img>
					<div class="card-item__big-description">
						<prismic-rich-text :htmlSerializer="bigDesc" :field="card.card_hidden_text" class="card-item__big-description--wrap" />
					</div>
				</v-card>
			</div>
		</div>
		<div class="card-list__two-col--col">
			<div class="card-item" v-for="(card, i) in evenCards" :key="i" @click="$set(card, 'selected', !card.selected)" :class="{ 'card-item--active': card.selected }">
				<v-card class="card-item__card" elevation="10" light>
					<font-awesome-layers class="fa-3x card-item__icon" v-if="card.icon">
						<font-awesome-icon class="card-item__icon--background" :icon="['fas', 'circle']" />
						<font-awesome-icon class="card-item__icon--plus" :icon="['fas', 'plus']" transform="shrink-6" />
					</font-awesome-layers>
					<v-img class="card-item__background" dark :src="card.card_bg.url" :alt="card.card_bg.alt">
						<h2 class="card-item__title">{{ $prismic.asText(card.card_title) }}</h2>
						<prismic-rich-text :field="card.card_text" class="card-item__text" />
					</v-img>
					<div class="card-item__big-description">
						<prismic-rich-text :htmlSerializer="bigDesc" :field="card.card_hidden_text" class="card-item__big-description--wrap" />
					</div>
				</v-card>
			</div>
		</div>
		<prismic-rich-text class="card-list__copy" :field="slice.primary.copy" v-if="slice.primary.copy.length > 0" />
	</div>
</template>
<script>
import '@fortawesome/fontawesome-svg-core/styles.css'
import prismicDOM from 'prismic-dom';
import linkResolver from '~/plugins/link-resolver.js';
const Elements = prismicDOM.RichText.Elements;

const cardText = function(type, element, content, children) {
	if (type === Elements.paragraph) {
		const wrapperClassList = [element.label || '']
		let result = `<p class="${wrapperClassList.join(' ')}">${content}</p>`;
		return result
	}
	if (type === Elements.hyperlink) {
		let result = '';
    const url = prismicDOM.Link.url(element.data, linkResolver);

    if (element.data.link_type === 'Document') {
      result = `<nuxt-link to="${url}">${content}</nuxt-link>`;
    } else {
			const target = element.data.target ? `target="'${element.data.target}'" rel="noopener"` : '';
      result = `<a class="${element.label}" href="${url}" ${target}>${content}</a>`;
    }
	}
	return null;
}
const bigDesc = function(type, element, content, children) {
	if (type === Elements.hyperlink) {
		let result = '';
    const url = prismicDOM.Link.url(element.data, linkResolver);

    if (element.data.link_type === 'Document') {
			result = `<router-link to="${url}">${content}</router-link>`;
			
    } else {
			const target = element.data.target ? `target="'${element.data.target}'" rel="noopener"` : '';
			result = `<a class="${element.label}" href="${url}" ${target}>${content}</a>`;
		}
	}
	return null;
}
export default {
	name: "TwoColCards",
	props: ["slice"],
	data() {
		return {
			isActive: "",
			selected: false,
			cardText,
			bigDesc
		}
	},
	computed: {
		evenCards() {
			var cards = this.slice.items
			const evenCards = []
			for (var i = 0; i < cards.length; i++) {
				if ((i + 1) % 2 === 0) {
					evenCards.push(cards[i])
				}
			}
			return evenCards
		},
		oddCards() {
			var cards = this.slice.items
			const oddCards = []
			for (var i = 0; i < cards.length; i++) {
				if ((i + 1) % 2 !== 0) {
					oddCards.push(cards[i])
				}
			}
			return oddCards
		},
	},
	methods: {
		serviceSelect: function(event) {
			this.isActive = !this.isActive
			if (event) {
				if (event.target.classList.length <= 0 || event.target.nodeName == "svg") {
					var elem = document.getElementsByTagName(event.target.tagName)[0]
				} else {
					elem = document.querySelector(`.${event.target.className}`)
				}
				var closest = elem.closest(".card-item__card")
				closest.classList.toggle("card-item__card--active")
			}
		},
	},
}
</script>
<style lang="scss">
.card-list {
	max-width: 150rem;
	margin: auto;
	&__two-col {
		padding: 0 2vw;
		position:relative;
		z-index:2;
		@include respond(mobileLarge) {
			padding:0 30px;
			display: grid;
			grid-template-columns: 50% 50%;
		}
	}
	&__copy {
		width:100%;
		grid-column:span 2;
		padding-bottom:10px;

		h2 {
			text-align:center;
			margin:10px 0;
		}

		h2, strong {
			color:$primary-dark;
		}
	}
	
}
.card-item {
	padding: 12px;
	cursor: pointer;
	&__background {
		height: 225px;
		padding: 0 2rem;
		display: flex;
		align-items: flex-end;
		z-index: 3 !important;
		@include respond(mobileLarge) {
			height: 300px;
		}
	}
	&__text {
		display:inline-flex;
		width:100%;
		
		&--transparent-bg {
			background-color:rgba($color-white, .7);
			font-size:1.5em;
			color:$color-black;
			font-family:$heading-font;
			padding:5px 20px;
		}
	}
	&__big-description {
		opacity: 0;
		height: 0;
		-webkit-transform: translateY(-154px);
		transform: translateY(-154px);
		-webkit-transition: all 0.3s ease-in-out;
		transition: all 0.3s ease-in-out;
		&--wrap {
			padding: 12px 20px 20px 20px;
			transform: translateY(0);
			background-color: transparent;
			margin-bottom: 0;
		}
		font-size:.9em;
		@include respond(tabletLarge) {
			font-size:1em;
		}
	}
	&__icon {
		justify-self: end;
		right: -15px;
		top: -16px;
		z-index: 4;
		position: absolute;
		backface-visibility: hidden;
		transform: rotate(0);
		transition: all 0.2s ease-in;
		&--background {
			color: $primary-dark;
		}

		&--plus {
			// transform: rotate(0) scale3d(1, 1, 1) translateZ(0);
			-webkit-transition: all 0.3s cubic-bezier(0.25, 0.8, 0.5, 1);
			transition: all 0.3s cubic-bezier(0.25, 0.8, 0.5, 1);
			color: $color-white;
			backface-visibility: hidden;
			-webkit-transform: scale(0.7);
			transform: scale(0.7);
			-webkit-font-smoothing: subpixel-antialiased;
		}
	}
	&:hover {
		.card-item__icon {
		}
	}
	&--active {
		.card-item__icon {
			transform: rotate(270deg);
			transition: all 0.2s ease-in;
		}
		.card-item__big-description {
			display: flex;
			height: 210px;
			opacity: 1;
			-webkit-transform: translateY(0);
			transform: translateY(0);
			-webkit-transition: all 0.6s cubic-bezier(0.67, 0.01, 0.52, 1.4);
			transition: all 0.6s cubic-bezier(0.67, 0.01, 0.52, 1.4);
			animation: none;
			@include respond(mobileLarge) {
				height: 190px;
				opacity: 1;
				-webkit-transform: translateY(0);
				-webkit-transition: all 0.6s cubic-bezier(0.67, 0.01, 0.52, 1.4);
				transition: all 0.6s cubic-bezier(0.67, 0.01, 0.52, 1.4);
				animation: none;
			}
			@include respond(tabletSmall) {
				height: 200px;
				opacity: 1;
				-webkit-transform: translateY(0);
				transform: translateY(0);
				-webkit-transition: all 0.6s cubic-bezier(0.67, 0.01, 0.52, 1.4);
				transition: all 0.6s cubic-bezier(0.67, 0.01, 0.52, 1.4);
				animation: none;
			}
		}
	}
}
.card-list__two-col--about-us {
	column-gap:10px;
	.card-list__two-col--col {
		&:nth-of-type(1) {
			.card-item {
				&__text {
					justify-content:flex-end;
				}
				&:nth-of-type(1) {
					.card-item__card {
						background:url('https://images.prismic.io/guardianrestoration/b14f9686-12b9-4c52-8971-0b0b05c7d6d2_Light+Rectangle+BG+for+G+About+Us+top+section.png?auto=compress,format') no-repeat center center!important;
					}
				}
				&:nth-of-type(2) {
					.card-item__card {
						background:url('https://images.prismic.io/guardianrestoration/9dd15327-1507-4499-af51-dd230fe5c5b1_Darker+Rectangle+BG+for+G+About+Us+top+section.png?auto=compress,format') no-repeat center center!important;
					}
				}
			}
		}
		&:nth-of-type(2) {
			.card-item {
				&__text {
					justify-content: flex-start;
				}
				&:nth-of-type(1) {
					.card-item__card {
						background:url('https://images.prismic.io/guardianrestoration/9dd15327-1507-4499-af51-dd230fe5c5b1_Darker+Rectangle+BG+for+G+About+Us+top+section.png?auto=compress,format') no-repeat center center!important;
					}
				}
				&:nth-of-type(2) {
					.card-item__card {
						background:url('https://images.prismic.io/guardianrestoration/b14f9686-12b9-4c52-8971-0b0b05c7d6d2_Light+Rectangle+BG+for+G+About+Us+top+section.png?auto=compress,format') no-repeat center center!important;
					}
				}
			}
		}
	}
	
	
	.card-item {
		border:5px solid $color-white;
		&__card {
			overflow:hidden;
		}
		&__big-description {
			background-color:rgb(224 215 213 / 45%);
			transform:rotate(0deg) scale(1);
			opacity:0;
			transition: all .3s ease-out;
			height:100%;
			position:absolute;
			top:0;
			width:100%;
			overflow:hidden;
			color:$color-black;

			&--wrap {
				transform:translateY(-340px);
				transition: all .2s ease-in-out;
				padding:12px 0px 20px 0px;
				text-align:center;

				p {					
					&:not(:last-child) {
						padding:20px 25px 5px 25px;
					}
				}
			}
		}

		&__hidden-text-heading {
			display:block;
			width:100%;
			background:$color-grey;
			color:$color-white;
		}

		&__background {
			opacity:1;
			transition: all .5s ease-out;
			height:300px;
			@include respond(mobileLarge) {
				height:350px;
			}
			@include respond(desktopSmall) {
				height:300px;
			}
		}

		// &:hover {
		// 	.card-item__background {
		// 		transform:rotate(720deg) scale(0);
		// 		opacity:0;
		// 	}
		// }

		&--active {
			.card-item__big-description {
				opacity:1;
				transform:translateY(0px) rotate(0deg);
				transition-delay:.4s;
			}
			.card-item__background {
				transform:rotate(720deg) scale(0);
				opacity:0;
			}
			.card-item__big-description--wrap {
				transform:translateY(0px);
				transition-delay: .5s;
			}
		}
	}
}
</style>
