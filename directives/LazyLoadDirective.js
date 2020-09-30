export default {
  inserted: el => {
    const config = {
      root: null,
      rootMargin: '0px',
      threshold: 0
    }

    const imageElement = Array.from(el.children).find(el => el.nodeName === "IMG");

    function preloadImages() {
      imageElement.src = imageElement.dataset.url
    }
    function createObserver() {
      const options = {
        root: null,
        threshold: "0"
      };
      const observer = new IntersectionObserver(handleIntersect, config);
      observer.observe(el);
    }

    function handleIntersect(entries, observer) {
      entries.forEach(entry => {
        if (entry.isIntersecting) {
          preloadImages();
          observer.unobserve(el)
        }
      });
    }

    if (window["IntersectionObserver"]) {
      createObserver();
    } else {
      preloadImages();
    }
  }
}