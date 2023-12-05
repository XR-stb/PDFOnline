import React, { useEffect, useState } from "react";
import {message, Row, Col, FloatButton, Modal, Popover} from "antd";
import { UploadOutlined } from '@ant-design/icons';
import {PdfType, UserType} from "../../types";
import { listPdfs } from "../../api/pdfonline/pdf";
import CardComponent from "../../components/card";
import UploadForm from "./upload";

interface PDFContentProps {
  user: {user: UserType | undefined, loggedIn: boolean, setLoggedIn: React.Dispatch<React.SetStateAction<boolean>>}
}

const PDFContent: React.FC<PDFContentProps> = ({ user }) => {
  const [pdfs, setPdfs] = useState<PdfType[]>([]);

  useEffect(() => {
    const fetchData = async () => {
      try {
        await listPdfs()
          .then((data) => {
            setPdfs(data);
          })
          .catch((error) => {
            message.error(error.message);
          });
      } catch (error) {
        console.log(error);
      }
    };

    fetchData();
  }, []);

  const totalItems = pdfs.length;
  const halfTotalItems = Math.ceil(totalItems / 2);
  const firstColumn = pdfs.slice(0, halfTotalItems);
  const secondColumn = pdfs.slice(halfTotalItems);

  const [isModalOpen, setIsModalOpen] = useState(false);

  const handleUpload = () => {
    setIsModalOpen(true);
  };

  const handleOk = () => {
    setIsModalOpen(false);
  };

  const handleCancel = () => {
    setIsModalOpen(false);
  };

  return (
    <>
      {
        user.loggedIn ?
          <FloatButton icon={<UploadOutlined />} onClick={handleUpload} /> :
          <Popover placement="topLeft" content={<>Please log in first</>}>
            <FloatButton icon={<UploadOutlined />} />
          </Popover>
      }
      <Modal title="Upload" footer={null} open={isModalOpen} onOk={handleOk} onCancel={handleCancel}>
        <UploadForm />
      </Modal>
      <Row gutter={[16, 16]}>
        <Col span={12}>
          {firstColumn.map((pdf: PdfType) => (
            <div key={pdf.id} style={{ marginBottom: "16px" }}>
              <CardComponent pdf={pdf} />
            </div>
          ))}
        </Col>
        <Col span={12}>
          {secondColumn.map((pdf: PdfType) => (
            <div key={pdf.id} style={{ marginBottom: "16px" }}>
              <CardComponent pdf={pdf} />
            </div>
          ))}
        </Col>
      </Row>
    </>
  );
};

export default PDFContent;
