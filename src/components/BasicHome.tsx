import Framily from "./assets/framily.jpg";
import "./App.css";
import { Box, Typography, Container } from "@mui/material";

function BasicHome() {
  return (
    <Container maxWidth="lg">
      <Box
        display="flex"
        justifyContent="center"
        alignItems="center"
        height="100%"
        width="100%"
        overflow="hidden"
        paddingBottom="10px"
      >
        <Box
          component="img"
          sx={{
            maxWidth: "500px",
            width: "90%",
            height: "auto",
            objectFit: "scale-down",
          }}
          alt="Centered Image"
          src={Framily}
        />
      </Box>
      <Typography textAlign="center" fontSize="20px">
        Welcome to Ian Myjer dot Com!
      </Typography>
    </Container>
  );
}

export default BasicHome;
