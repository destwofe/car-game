import Image from "next/image"

import carN from '../public/images/cars/car-n.png'
import carE from '../public/images/cars/car-e.png'
import carS from '../public/images/cars/car-s.png'
import carW from '../public/images/cars/car-w.png'

const GameBoard = (props: { dimensionX: number; dimensionY: number; carPositionX: number | undefined; carPositionY: number | undefined; carDirection: string | undefined }) => {
  const getCarImage = (carDirection: string | undefined) => {
    switch (carDirection) {
      case "NORTH":
        return carN
      case "EAST":
        return carE
      case "SOUTH":
        return carS
      case "WEST":
        return carW
      default:
        return carN
    }
  }

  const { dimensionX, dimensionY, carPositionX, carPositionY, carDirection } = props

  return (
    <div className="board">
      <div className="wrap">
        {Array(dimensionY).fill(1).map((_, i) => (
          <div key={i} className="x">{
            Array(dimensionX).fill(1).map((_, j) => (
              <div key={j} className="block">
                {
                  dimensionY - i === carPositionY && j + 1 === carPositionX ?
                    <Image src={getCarImage(carDirection)} alt="carN" width={100} height={100} /> :
                    `${dimensionY - i},${j + 1}`
                }
              </div>
            ))
          }</div>
        ))}
      </div>
    </div>
  )
}

export default GameBoard