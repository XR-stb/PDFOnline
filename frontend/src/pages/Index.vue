<template>
    <div id="app">
        <!-- 页眉 -->
        <header>
            <h1>PDF Online</h1>
        </header>

        <div class="maintainer">
            <div class="left">
                <div class="theme slider">
                    <img src="../assets/更多.png" alt="" class="slider-img">
                </div>
                <div class="login slider">
                    <router-link to="/login">
                        <img src="../assets/登录.png" alt="" class="slider-img">
                    </router-link>
                </div>
                <div class="upload slider" @click="openUploadDialog">
                    <img src="../assets/上传.png" alt="" class="slider-img">
                </div>
                <div class="upload-container" v-if="showUploadDialog">
                    <el-upload class="upload-demo upload-box" drag
                        action="https://run.mocky.io/v3/9d059bf9-4660-45f2-925d-ce80ad6c4d15" multiple>
                        <!-- <el-icon class="el-icon--upload"><upload-filled /></el-icon> -->
                        <img src="../assets/upload.png" alt="" class="upload-img">
                        <div class="el-upload__text upload-bg">
                            Drop file here or <em>click to upload</em>
                        </div>
                        <template #tip>
                            <div class="el-upload__tip">
                                jpg/png files with a size less than 500kb
                            </div>
                        </template>
                    </el-upload>
                </div>
                <div class="home slider">
                    <img src="../assets/首页.png" alt="" class="slider-img">
                </div>
            </div>
            <div class="right">
                <!-- PDF列表 -->
                <div class="pdf-container">
                    <div v-for="pdf in pdfList" :key="pdf.title" class="pdf-block" @click="openPDF(pdf.url)">
                        <!-- <img :src="pdf.cover_url" alt="加载封面失败" class="cover"> -->
                        <el-image :src="pdf.cover_url" :fit="fit" class="cover"></el-image>
                        <p class="title">{{ pdf.title }}</p>
                        <p class="desc">{{ pdf.description }}</p>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
  
<script>
import { UploadFilled } from '@element-plus/icons-vue'
export default {
    data() {
        return {
            pdfList: [], // 所有 PDF 数据
            displayedPDFs: [], // 当前页面显示的 PDF 数据
            itemsPerPage: 8, // 每页显示的 PDF 数量
            currentPage: 1, // 当前页数
            showUploadDialog: false
        };
    },
    created() {
        this.fetchPDFList();
    },
    methods: {
        openUploadDialog() {
            this.showUploadDialog= true
        },
        async fetchPDFList() {
            try {
                const response = await this.$axios.get(
                    "http://localhost:8080/v1/pdfs"
                );
                this.pdfList = response.data.pdfs;
                this.downloadBaseURL = "http://localhost:8080";
                console.log("pdfList:", this.pdfList);
            } catch (error) {
                console.error("Error fetching PDF list:", error);
            }
        },
        openPDF(pdfFileName) {
            console.log("filename:" + pdfFileName);
            // const fileName = pdfFileName.substring(
            //     pdfFileName.lastIndexOf("/") + 1
            // );
            const pdfUrl = this.downloadBaseURL + pdfFileName;
            window.location.href = pdfUrl;
        },
        updateDisplayedPDFs() {
            const startIdx = (this.currentPage - 1) * this.itemsPerPage;
            const endIdx = startIdx + this.itemsPerPage;
            this.displayedPDFs = this.pdfList.slice(startIdx, endIdx);
        },
    },
};
</script>

<style>
@import "../css/index.css";
</style>

  