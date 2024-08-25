import Framily from "./assets/framily.jpg";
import "./App.css";
import { Stack, Grid, Paper, Container, Typography } from "@mui/material";

function App() {
  return (
    <Container maxWidth="lg" sx={{ pt: 10 }}>
      <Grid container spacing={2}>
        <Grid item xs={12} md={4}>
          <Paper sx={{ p: 2, height: "100%" }}>
            <Stack spacing={2}>
              <Typography>Item 1</Typography>
              <Typography>Item 2</Typography>
              <Typography>Item 3</Typography>
            </Stack>
          </Paper>
        </Grid>
        <Grid item xs={12} md={8}>
          <Paper sx={{ p: 2, height: "100%" }}>Right Column (2/3)</Paper>
        </Grid>
      </Grid>
    </Container>
  );
}

export default App;
