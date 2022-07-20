import { Autocomplete, Button, TextField } from "@mui/material"
import { useState } from "react"

const PlaceCarForm = (props: { onPlaceHandler: (carPositionX: number, carPositionY: number, carDirection: string) => void }) => {
  const [carPositionX, setCarPositionX] = useState<number>(3)
  const [carPositionY, setCarPositionY] = useState<number>(3)
  const [carDirection, setCarDirection] = useState<string>('WEST')

  return (
    <form className="form">
      <div className="line">
        <TextField label="Position X" value={carPositionX} onChange={(e) => setCarPositionX(Number(e.target.value))} type="number" />
        <TextField label="Position Y" value={carPositionY} onChange={(e) => setCarPositionY(Number(e.target.value))} type="number" />
      </div>
      <Autocomplete
      className="line"
        value={carDirection}
        onChange={(e) => setCarDirection(e.currentTarget.innerHTML)}
        disablePortal
        options={["NORTH", "EASE", "SOUTH", "WEST"]}
        sx={{ width: 200 }}
        renderInput={(params) => (
          <TextField key={params.id} {...params} label="Direction" />
        )} />
      <Button variant="contained" onClick={() => {
        if (carPositionX == null) return
        if (carPositionY == null) return
        if (carDirection == null) return
        props.onPlaceHandler(carPositionX, carPositionY, carDirection)
      }}>Place Car</Button>
    </form>
  )
}

export default PlaceCarForm
