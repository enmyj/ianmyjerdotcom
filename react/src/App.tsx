import Framily from "./assets/framily.jpg";
import "./App.css";
import { Box, Typography } from "@mui/material";

function App() {
  return (
    <Box max-width="1280px" padding="10px" textAlign="center">
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
      <Typography fontSize="20px">Welcome to Ian Myjer dot Com!</Typography>
    </Box>
  );
}

export default App;
