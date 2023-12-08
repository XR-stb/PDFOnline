import React, { useState } from "react";
import { Modal } from "antd";
import ResetPasswordForm from "../form/resetPass";

const useResetPasswordModal = () => {
  const [isResetPasswordModalOpen, setResetPasswordModalOpen] = useState(false)
  const showResetPasswordModal = () => setResetPasswordModalOpen(true)
  const hideResetPasswordModal = () => setResetPasswordModalOpen(false)

  const resetPasswordModal = (
    <Modal title="Reset Password" footer={null} open={isResetPasswordModalOpen} onCancel={hideResetPasswordModal}>
      <ResetPasswordForm hideResetPasswordModal={hideResetPasswordModal} />
    </Modal>
  )

  return { resetPasswordModal, showResetPasswordModal };
}

export default useResetPasswordModal;