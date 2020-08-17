<template>
	<div class="card-list__two-col card-list">
		<div class="card-list__two-col--col">
			<div class="card-item" v-for="(card, i) in oddCards" :key="i" @click="$set(card, 'selected', !card.selected)" :class="{ 'card-item--active': card.selected }">
				<v-card class="card-item__card" elevation="10" light>
					<font-awesome-layers class="fa-3x card-item__icon">
						<font-awesome-icon class="card-item__icon--background" :icon="['fas', 'circle']" />
						<font-awesome-icon class="card-item__icon--plus" :icon="['fas', 'plus']" transform="shrink-6" />
					</font-awesome-layers>
					<v-img class="card-item__background" dark :src="card.card_bg.url" :alt="card.card_bg.alt" gradient="120deg, rgba(0,0,0,0.88) 0%, rgba(0,0,0,0.6) 45%, rgba(0,0,0,0.4) 60%,  rgba(0,0,0,0) 100%">
						<h2 class="card-item__title">{{ $prismic.asText(card.card_title) }}</h2>
						<p class="card-item__text">{{ $prismic.asText(card.card_text) }}</p>
					</v-img>
					<div class="card-item__big-description">
						<div class="card-item__big-description--wrap">
							<p class="card-item__hidden-text">{{ $prismic.asText(card.card_hidden_text) }}</p>
						</div>
					</div>
				</v-card>
			</div>
		</div>
		<div class="card-list__two-col--col">
			<div class="card-item" v-for="(card, i) in evenCards" :key="i" @click="$set(card, 'selected', !card.selected)" :class="{ 'card-item--active': card.selected }">
				<v-card class="card-item__card" elevation="10" light>
					<font-awesome-layers class="fa-3x card-item__icon">
						<font-awesome-icon class="card-item__icon--background" :icon="['fas', 'circle']" />
						<font-awesome-icon class="card-item__icon--plus" :icon="['fas', 'plus']" transform="shrink-6" />
					</font-awesome-layers>
					<v-img class="card-item__background" dark :src="card.card_bg.url" :alt="card.card_bg.alt" gradient="120deg, rgba(0,0,0,0.88) 0%, rgba(0,0,0,0.6) 45%, rgba(0,0,0,0.4) 60%,  rgba(0,0,0,0) 100%">
						<h2 class="card-item__title">{{ $prismic.asText(card.card_title) }}</h2>
						<prismic-rich-text :field="card.card_text" />
					</v-img>
					<div class="card-item__big-description">
						<prismic-rich-text :field="card.card_hidden_text" class="card-item__big-description--wrap" />
					</div>
				</v-card>
			</div>
		</div>
	</div>
</template>
<script>
import { library } from "@fortawesome/fontawesome-svg-core"
import { faPlus, faCircle } from "@fortawesome/free-solid-svg-icons"
import { FontAwesomeIcon, FontAwesomeLayers } from "@fortawesome/vue-fontawesome"
library.add(faPlus, faCircle)
export default {
	name: "TwoColCards",
	props: ["slice"],
	components: {
		FontAwesomeIcon,
		FontAwesomeLayers,
	},
	data() {
		return {
			isActive: "",
			selected: undefined,
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
		display: grid;
		@include respond(mobileLarge) {
			grid-template-columns: 50% 50%;
		}
	}
}
.card-item {
	padding: 12px;
	cursor: pointer;
	&__background {
		height: 300px;
		padding: 0 2rem;
		display: flex;
		align-items: flex-end;
		z-index: 3 !important;
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
			height: 187px;
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
</style>
