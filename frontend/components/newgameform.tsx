import { Button, TextField } from "@mui/material"
import { useState } from "react"

const NewGameForm = (props: { onNewGameHandler: (carDimensionX: number, carDimensionY: number) => void }) => {
  const [carDimensionX, setCarDimensionX] = useState<number>(5)
  const [carDimensionY, setCarDimensionY] = useState<number>(5)

  return (
    <form className="form">
      <div className="line">
        <TextField label="Dimension X" value={carDimensionX} onChange={(e) => setCarDimensionX(Number(e.target.value))} type="number" />
        <TextField label="Dimension Y" value={carDimensionX} onChange={(e) => setCarDimensionY(Number(e.target.value))} type="number" />
      </div>
      <Button variant="contained" onClick={() => {
        if (carDimensionX == null) return
        if (carDimensionY == null) return
        props.onNewGameHandler(carDimensionX, carDimensionY)
      }}>New Game</Button>
    </form>
  )
}

export default NewGameForm
