type Props = { src: string; type?: string }
export default function VideoPlayer({ src, type = 'video/mp4' }: Props) {
  return (
    <video className="video" controls preload="metadata">
      <source src={src} type={type} />
      Your browser does not support the video tag.
    </video>
  )
}
