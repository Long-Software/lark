type TimerContainerProps = {
  title: string
  time: number
}

export const TimerContainer = ({ title, time }: TimerContainerProps) => {
  return (
    <div className='w-[230px]'>
      <p className='text-center text-4xl'>{title}</p>
      <div className='mt-3 rounded-lg border border-foreground p-24 text-5xl'>{time}</div>
    </div>
  )
}

export default TimerContainer
