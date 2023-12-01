import { useState} from "react";
import { Button, Card, Flex, Typography } from "antd";
import { LoadingOutlined } from "@ant-design/icons";
import { cardStyle, imgStyle } from "./styles";
import { Pdf } from "../../api/pdfonline/pdf"

interface CardComponentProps {
  pdf: Pdf
}

const CardComponent = ({pdf}: CardComponentProps) => {
  const [onLoad, setOnLoad] = useState(false)
  const handleOnLoad = () => {
    setOnLoad(true)
  }

  return (
    <Card hoverable style={cardStyle} bodyStyle={{ padding: 0, overflow: 'hidden' }}>
      <Flex justify="space-between">
        {onLoad ? (<></>) : (<LoadingOutlined />)}
        <img alt="avatar" src={pdf.url} style={imgStyle} onLoad={handleOnLoad} />
        <Flex vertical align="flex-end" justify="space-between" style={{ padding: 32 }}>
          <Typography.Title level={3}>
            {pdf.title}
          </Typography.Title>
          <Typography.Text>
            {pdf.description}
          </Typography.Text>
          <Button type="primary" href={pdf.url} target="_blank">
            Get Start
          </Button>
        </Flex>
      </Flex>
    </Card>
  );
}

export default CardComponent;