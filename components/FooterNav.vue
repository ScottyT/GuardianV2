<template>
  <div class="footer-wrapper darken-4">
    <ul class="footer-wrapper__items">
      <li v-for="(menuLink, index) in $store.state.footer.menu_links" :key="index" class="footer-wrapper__menu-item">
        <img class="footer-wrapper__item-image" :src="menuLink.link_image.url" />
        <prismic-link :field="menuLink.link">{{ $prismic.asText(menuLink.label) }}</prismic-link>
      </li>
    </ul>
		<div class="footer-wrapper__bottom">
			<img src="https://images.prismic.io/guardianrestoration/63a35591-7f6b-4f3e-9600-bdb13206e5e9_Guardian+Logo+text+black+transp.png?auto=compress,format" class="footer-wrapper__bottom--logo" />
			<div class="footer-wrapper__social-icons">
				<div v-for="n in 2" :key="n">
					<p>Icon</p>
				</div>
			</div>
		</div>
  </div>
</template>
<script>
  export default {
    name: "FooterNav",
  }
</script>
<style lang="scss">
  .footer-wrapper {
    background-image: url('https://images.prismic.io/guardianrestoration/1810cd43-a5b6-4590-88c6-ff6e05bb5392_Footer+BG+image.png?auto=compress,format');
    background-size: cover;
    padding: 3.7rem 0 2rem 0;
		height: 420px;
		position:relative;
		display:grid;
		grid-template-rows:repeat(2, 1fr);

    &__items {
      flex-direction: row;
      justify-content: space-between;
      width: 100%;
      padding-left: 0 !important;
      max-width: 1200px;
      margin: auto;

      @include respond(mobileLarge) {
        display: flex;
      }
		}
		
		&__social-icons {
			display:flex;
			justify-content:flex-end;
			width:100%;
			max-width:500px;
		}

    &__item-image {
      max-width: 100px;
      object-fit: contain;
      height: 69px;
    }

    &__menu-item {
      padding: 0 4px;
      line-height: 20px;

      a {
        display: flex;
        padding: 0 20px;
        background-image: linear-gradient(120deg, transparent 0%, transparent 50%, rgba($color-white, 0.2) 50%);
        background-size: 245%;
        transition: all 0.4s ease-in;
        color: inherit !important;
        font-weight: 500;
        text-transform: uppercase;

        &:hover {
          background-position: 100%;
        }
      }

      &--services {}
		}
		
		&__bottom {
			display:flex;
			align-items:center;
			max-width:1200px;
			margin:0 auto;
			justify-content:space-between;
			width:100%;
			&--logo {
				object-fit:contain;
				height:50px;
				width:200px;
			}
		}
  }

  .footer {
    margin: auto;
    max-width: 114rem;
    padding: 1.7rem 0 2rem 0;

    #services-list {
      padding-left: 15px;
      padding-top: 7px;
    }

    &__site-title {
      width: 100%;
      text-align: center;
      margin-bottom: 1.5rem;
    }


  }
</style>