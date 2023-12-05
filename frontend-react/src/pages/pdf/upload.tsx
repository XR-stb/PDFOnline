import {Button, Form, message} from "antd";
import {DescriptionItem, TitleItem, UploadDraggerItem} from "../../components/form/items";

const UploadForm = () => {
  const [form] = Form.useForm();
  const onFinish = () =>  {

  };

  const onInputChange = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
    const value = e.target.value.replace(/\n/g, '');
    form.setFieldsValue({ description: value });
  };

  return (
    <Form
      form={form}
      name="upload"
      onFinish={onFinish}
      scrollToFirstError
    >
      <TitleItem />
      <DescriptionItem onChange={onInputChange} />
      <UploadDraggerItem />
      <Button type="primary" htmlType="submit">
        Upload
      </Button>
    </Form>
  );
}

export default UploadForm;