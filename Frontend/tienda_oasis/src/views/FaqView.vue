<template>
  
  <div class="faq-container">
    <h2> Preguntas Frecuentes </h2>
    <Faq 
    v-for="(item, index) in faqList"
    :key="index"
    :question="item.question"
    :answer="item.answer"
    />
  </div>
</template>

<script>
    import Faq from '@/components/Faq.vue'
import NavBar from '@/components/menus/NavBar.vue';
    export default {
        name: 'FaqView',
        components: {
            Faq
        },
        data() {
            return {
                faqList: [
                    {
                        question: '¿Qué es un tienda virtual?',
                        answer: 'Una tienda virtual es un espacio virtual donde los clientes pueden comprar productos y servicios de manera segura y confiable.'
                    },
                    {
                        question: '¿Qué es una tienda virtual de Amazon?',
                        answer: 'Una tienda virtual de Amazon es un espacio virtual donde los clientes pueden comprar productos y servicios de manera segura y confiable.'
                    },
                    {
                        question: '¿Qué es una tienda virtual de Netflix?',
                        answer: 'Una tienda virtual de Netflix es un espacio virtual donde los clientes pueden comprar productos y servicios de manera segura y confiable.'
                    },
                    {
                        question: '¿Qué es una tienda virtual de Youtube?',
                        answer: 'Una tienda virtual de Youtube es un espacio virtual donde los clientes pueden comprar productos y servicios de manera segura y confiable.'
                    },
                    {
                        question: '¿Qué es una tienda virtual de Twitch?',
                        answer: 'Una tienda virtual de Twitch es un espacio virtual donde los clientes pueden comprar productos y servicios de manera segura y confiable.'
                    },
                    {
                        question: '¿Qué es una tienda virtual de Disney?',
                        answer: 'Una tienda virtual de Disney es un espacio virtual donde los clientes pueden comprar productos y servicios de manera segura y confiable.'
                    },
                    {
                        question: '¿Qué es una tienda virtual de Nintendo?',
                        answer: 'Una tienda virtual de Nintendo es un espacio virtual donde los clientes pueden comprar productos y servicios de manera segura y confiable.'
                    },
                    {
                        question: '¿Qué es una tienda virtual de Google?',
                        answer: 'Una tienda virtual de Google es un espacio virtual donde los clientes pueden comprar productos y servicios de manera segura y confiable.'
                    },
                    {
                        question: '¿Qué es una tienda virtual de Microsoft?',
                        answer: 'Una tienda virtual de Microsoft es un espacio virtual donde los clientes pueden comprar productos y servicios de manera segura y confiable.'
                    }]
                }
            }
        }
    
</script>
    
<style scoped>

    .faq-container{
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        height: 100vh;
        text-align: center;
    }

    h2{
        margin-bottom: 20px;
        color: #007BFF;
    }

</style>