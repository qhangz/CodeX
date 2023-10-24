<template>
  <div class="about">
    <h1>This is an about page </h1>
    <RouterLink to="/">Home</RouterLink>
  </div>
</template>

<style lang="scss" scoped>
// 页面居中显示
.about {
  width: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
}


@media (min-width: 1024px) {
  .about {
    // min-height: 100vh;
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
  }
}
</style>
