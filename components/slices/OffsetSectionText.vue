<template>
	<div :class="`offset-section-text ${slice.primary.label !== undefined ? slice.primary.label : ''}`">
		<div class="offset-section-text__row" v-for="(section, i) in slice.items" :key="i">
			<div :class="`offset-section-text__content offset-section-text__content--${(i + 1) % 2 === 0 ? 'left' : 'right'}`">
				<prismic-rich-text class="block-heading" :htmlSerializer="heading" :field="section.sectionheading" v-if="section.sectionheading.length > 0" />
				<prismic-rich-text :htmlSerializer="mainText" :field="section.sectiontext" />
			</div>
			
			<lazy-image v-if="$route.name === 'index'" :source="section.sectionimage.url" :alt="section.sectionimage.alt" :class="`offset-section-text__image offset-section-text__image--${(i + 1) % 2 === 0 ? 'left' : 'right'}`" />
			<push-slide v-if="slice.primary.label == 'about-us'" :image="section.sectionimage" :element="'offset-section-text__image'" :hidden="section.hidden_text"/>
		</div>
	</div>
</template>
<script>
import prismicDOM from "prismic-dom";
import linkResolver from '~/plugins/link-resolver.js';
const Elements = prismicDOM.RichText.Elements;

const mainText = function(type, element, content, children) {
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
}

const heading = function (type, element, content, children) {
	if (type === Elements.image) {
		var result = `<img src="${element.url}" alt="${element.alt || ''}" />`
		const wrapperClassList = [element.label || '', 'block-img'];
		result = `<div class="${wrapperClassList.join(' ')}">${result}</div>`
		return result
	}
	
	switch (type) {
		case Elements.paragraph:
			return `<p class="offset-section-text__content--subheading">${children.join("")}</p>`;
		// case Elements.image: 
		// 	const wrapperClassList = [element.label || '', 'block-img'];
		// 	const result = `<img src="${element.url}" alt="${element.alt || ''}" class="${wrapperClassList.join(' ')} />`;
		// 	return result;
		default:
			return null;
	}
	return null;
};
export default {
	name: "OffsetSectionText",
	props: ["slice"],
	data() {
		return {
			heading,
			mainText
		};
	},
};
</script>
<style lang="scss">
.offset-section-text {
	.block-img img {
	}
	
	&.about-us {
		padding-top:40px;
		.block-heading h2 {
			text-align:center;
			color:$primary;
		}
	}
	&__small-img {
		max-width:200px;
		margin:auto;
	}
	&__heading {
		&--army-rust {
			font-family:"Army Rust";
			font-size:2em;
			letter-spacing:12px;
			color:$primary;
			position:relative;

			&::before, &::after {
				content:'';
				position:absolute;
				width:100%;
				height:15px;
				left:0;
				background:url('https://images.prismic.io/guardianrestoration/6f3bc3ec-1109-4f93-9e5b-8284ad5b9925_Red+Line+for+COMPLETE.png?auto=compress,format');
				background-size:contain;
				background-repeat:no-repeat;

				@include respond(mobileLargeMax) {
					height:10px;
				}
			}
			&::before {
				top:-5px;
			}
			&::after {
				bottom:-5px;
			}
		}
		&--light {
			font-weight: 300;
		}
	}
	&__row {
		display: flex;
		flex-wrap: wrap;
		flex-direction:column;
		@include respond(mobileLarge) {
			padding:0 0 3rem;
			flex-direction:row;
		}
		@include respond(tabletLarge) {
			padding: 0px 0 6rem 0;
		}

		&:first-child {
			.offset-section-text__content {
				margin-top: 0rem;
				position: relative;
			}
		}
		&:nth-of-type(even) {
			@include respond(mobileLarge) {
				flex-direction: row-reverse;
			}
			.push-slide-wrapper {
				&__button {
					right:-100px;
					left:unset;
				}
			}
		}
	}
	&__content {
		display: grid;
		place-content: center;
		row-gap:20px;
		grid-template-columns:auto;
		padding:0 10px;
		
		@include respond(mobileLarge) {
			width: 50%;
		}
		@include respond(desktopSmall) {
			grid-template-columns: 500px;
		}
		h2 {
			text-transform:uppercase;
			font-weight:700;
		}
		&--left {
			padding: 32px;
			justify-content:flex-start;
			padding-left:50px;
		}
		&--right {
			padding: 0 28px 0 40px;
		}
		&--subheading {
			font-size: 1.2em;
			text-transform: uppercase;
		}
		&--small-image {
			height:25px;
			right:unset;
			left:139px;		
			position:absolute;
			bottom:12px;
			
			@include respond(mobileLarge) {
				width:133px;
				right:33px;
				left:unset;
			}
			@include respond(desktopSmall) {
				width:179px;
				height:40px;
			}
			img {
				object-fit:contain;
			}
		}
	}
	
	&__image {
		width:100%;
		position: relative;
		margin:auto;

		@include respond(mobileSmallPort) {
			width: 50%;
		}

		&--right {
			max-width: 690px;
			min-height: 250px;
			
			@include respond(mobileLarge) {
				left:-20px;
			}
		}
		&--left {
			max-width: 690px;
			//max-height: 333px;
			@include respond(mobileLarge) {
				right:-10px;
			}
		}
	}
}
</style>
