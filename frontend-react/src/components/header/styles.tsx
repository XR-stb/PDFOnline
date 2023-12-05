import React from "react";

export const headerStyle:React.CSSProperties = {
  position: 'sticky',
  top: 0,
  zIndex: 1,
  width: '100%',
  display: 'flex',
  backgroundColor: '#fff',
  alignItems: 'center',
  borderBottom: '1px solid #ebedf0',
  justifyContent: 'space-between',
  padding: '0 16px',
}

export const userPopoverStyle:React.CSSProperties = {
  maxWidth: '80%',
  margin: 'auto',
}

export const popoverTitleStyle:React.CSSProperties = {
  borderBottom: '1px solid #ebedf0',
  fontSize: '20px',
  fontWeight: 'bold',
}

export const logoutIconStyle:React.CSSProperties = {
  color: 'grey',
}