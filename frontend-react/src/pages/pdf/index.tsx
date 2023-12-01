import React, { useEffect, useState } from "react";
import { message, Row, Col } from "antd";
import BasicContainer from "../../containers/basic";
import { listPdfs, Pdf } from "../../api/pdfonline/pdf";
import CardComponent from "../../components/card";

const PDF: React.FC = () => {
  const [pdfs, setPdfs] = useState<Pdf[]>([]);

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

  return (
    <BasicContainer>
      <Row gutter={[16, 16]}>
        <Col span={12}>
          {firstColumn.map((pdf: Pdf) => (
            <div key={pdf.id} style={{ marginBottom: "16px" }}>
              <CardComponent pdf={pdf} />
            </div>
          ))}
        </Col>
        <Col span={12}>
          {secondColumn.map((pdf: Pdf) => (
            <div key={pdf.id} style={{ marginBottom: "16px" }}>
              <CardComponent pdf={pdf} />
            </div>
          ))}
        </Col>
      </Row>
    </BasicContainer>
  );
};

export default PDF;
